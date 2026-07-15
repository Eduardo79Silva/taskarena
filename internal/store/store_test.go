package store

import (
	"os"
	"testing"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/internal/testutil"
)

func withTempStoragePaths(t *testing.T) {
	t.Helper()

	dir := t.TempDir()

	origTasks := TasksFilePath
	origCurrent := CurrentTaskFilePath
	origCompleted := CompletedTasksFilePath

	TasksFilePath = dir + "/tasks.json"
	CurrentTaskFilePath = dir + "/current.json"
	CompletedTasksFilePath = dir + "/completed.json"

	t.Cleanup(func() {
		TasksFilePath = origTasks
		CurrentTaskFilePath = origCurrent
		CompletedTasksFilePath = origCompleted
	})
}

func TestReadTasksFile_MissingFileReturnsEmptySlice(t *testing.T) {
	withTempStoragePaths(t)

	tasks, err := ReadTasksFile(TasksFilePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("got %d tasks, want 0", len(tasks))
	}
}

func TestWriteReadTasksFile_RoundTrip(t *testing.T) {
	withTempStoragePaths(t)

	want := []task.Task{
		testutil.MakeTask("1", priority.High, 25, "work"),
		testutil.MakeTask("2", priority.Low, 10, "home"),
	}

	if err := WriteTasksFile(TasksFilePath, want); err != nil {
		t.Fatalf("writeTasksFile: %v", err)
	}

	got, err := ReadTasksFile(TasksFilePath)
	if err != nil {
		t.Fatalf("readTasksFile: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("got %d tasks, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].ID != want[i].ID || got[i].Name != want[i].Name {
			t.Errorf("task %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestReadTaskFile_MissingFileReturnsError(t *testing.T) {
	withTempStoragePaths(t)

	_, err := ReadTaskFile(CurrentTaskFilePath)
	if err == nil {
		t.Fatal("expected an error for a missing file, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected an os.IsNotExist error, got %v", err)
	}
}

func TestWriteReadTaskFile_RoundTrip(t *testing.T) {
	withTempStoragePaths(t)

	want := testutil.MakeTask("1", priority.High, 25, "work")

	if err := WriteTaskFile(CurrentTaskFilePath, want); err != nil {
		t.Fatalf("writeTaskFile: %v", err)
	}

	got, err := ReadTaskFile(CurrentTaskFilePath)
	if err != nil {
		t.Fatalf("readTaskFile: %v", err)
	}
	if got.ID != want.ID || got.Name != want.Name {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestPushTask_AppendsToExistingFile(t *testing.T) {
	withTempStoragePaths(t)

	first := testutil.MakeTask("1", priority.Medium, 25, "")
	second := testutil.MakeTask("2", priority.Medium, 25, "")

	PushTask(TasksFilePath, first)
	PushTask(TasksFilePath, second)

	tasks, err := ReadTasksFile(TasksFilePath)
	if err != nil {
		t.Fatalf("readTasksFile: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("got %d tasks, want 2", len(tasks))
	}
}

func TestWriteAllTasks_AndLoadTasks_RoundTrip(t *testing.T) {
	withTempStoragePaths(t)

	want := []task.Task{testutil.MakeTask("1", priority.Medium, 25, "")}
	WriteAllTasks(want)

	got := LoadTasks()

	if len(got) != 1 || got[0].ID != want[0].ID {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestDeleteTaskFromFile_RemovesTaskAndPersists(t *testing.T) {
	withTempStoragePaths(t)

	WriteAllTasks([]task.Task{
		testutil.MakeTask("1", priority.Medium, 25, ""),
		testutil.MakeTask("2", priority.Medium, 25, ""),
	})

	DeleteTaskFromFile(TasksFilePath, "1")

	got := LoadTasks()
	if len(got) != 1 || got[0].ID != "2" {
		t.Errorf("got %+v, want only task 2 remaining", got)
	}
}

func TestDeleteTask_RemovesMatchingID(t *testing.T) {
	tasks := []task.Task{
		testutil.MakeTask("1", priority.Medium, 25, ""),
		testutil.MakeTask("2", priority.Medium, 25, ""),
		testutil.MakeTask("3", priority.Medium, 25, ""),
	}

	got := DeleteTask(tasks, "2")

	if len(got) != 2 {
		t.Fatalf("got %d tasks, want 2", len(got))
	}
	for _, task := range got {
		if task.ID == "2" {
			t.Error("task with ID 2 was not removed")
		}
	}
}

func TestDeleteTask_UnknownIDIsNoOp(t *testing.T) {
	tasks := []task.Task{testutil.MakeTask("1", priority.Medium, 25, "")}

	got := DeleteTask(tasks, "does-not-exist")

	if len(got) != 1 {
		t.Errorf("got %d tasks, want 1 (unchanged)", len(got))
	}
}

func TestCompleteCurrentTask_MovesTaskToCompletedAndClearsCurrent(t *testing.T) {
	withTempStoragePaths(t)

	current := testutil.MakeTask("1", priority.High, 25, "work")
	WriteCurrentTask(current)

	CompleteCurrentTask()

	if _, err := os.Stat(CurrentTaskFilePath); !os.IsNotExist(err) {
		t.Errorf("expected current task file to be removed, stat err = %v", err)
	}

	completed, err := ReadTasksFile(CompletedTasksFilePath)
	if err != nil {
		t.Fatalf("readTasksFile(completed): %v", err)
	}
	if len(completed) != 1 {
		t.Fatalf("got %d completed tasks, want 1", len(completed))
	}
	if completed[0].ID != current.ID {
		t.Errorf("completed task ID = %q, want %q", completed[0].ID, current.ID)
	}
	if completed[0].CompletedAt == nil {
		t.Error("expected CompletedAt to be set, got nil")
	}
}
