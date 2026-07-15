package store

import (
	"os"
	"testing"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/internal/testutil"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()
	dir := t.TempDir()
	return &Store{
		TasksFilePath:          dir + "/tasks.json",
		CurrentTaskFilePath:    dir + "/current.json",
		CompletedTasksFilePath: dir + "/completed.json",
	}
}

func TestReadTasksFile_MissingFileReturnsEmptySlice(t *testing.T) {
	s := newTestStore(t)

	tasks, err := readTasksFile(s.TasksFilePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("got %d tasks, want 0", len(tasks))
	}
}

func TestWriteReadTasksFile_RoundTrip(t *testing.T) {
	s := newTestStore(t)

	want := []task.Task{
		testutil.MakeTask("1", priority.High, 25, "work"),
		testutil.MakeTask("2", priority.Low, 10, "home"),
	}

	if err := writeTasksFile(s.TasksFilePath, want); err != nil {
		t.Fatalf("writeTasksFile: %v", err)
	}

	got, err := readTasksFile(s.TasksFilePath)
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
	s := newTestStore(t)

	_, err := readTaskFile(s.CurrentTaskFilePath)
	if err == nil {
		t.Fatal("expected an error for a missing file, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected an os.IsNotExist error, got %v", err)
	}
}

func TestWriteReadTaskFile_RoundTrip(t *testing.T) {
	s := newTestStore(t)

	want := testutil.MakeTask("1", priority.High, 25, "work")

	if err := writeTaskFile(s.CurrentTaskFilePath, want); err != nil {
		t.Fatalf("writeTaskFile: %v", err)
	}

	got, err := readTaskFile(s.CurrentTaskFilePath)
	if err != nil {
		t.Fatalf("readTaskFile: %v", err)
	}
	if got.ID != want.ID || got.Name != want.Name {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestPushTask_AppendsToExistingFile(t *testing.T) {
	s := newTestStore(t)

	first := testutil.MakeTask("1", priority.Medium, 25, "")
	second := testutil.MakeTask("2", priority.Medium, 25, "")

	if err := s.PushTask(first); err != nil {
		t.Fatalf("PushTask: %v", err)
	}
	if err := s.PushTask(second); err != nil {
		t.Fatalf("PushTask: %v", err)
	}

	tasks, err := readTasksFile(s.TasksFilePath)
	if err != nil {
		t.Fatalf("readTasksFile: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("got %d tasks, want 2", len(tasks))
	}
}

func TestWriteAllTasks_AndLoadTasks_RoundTrip(t *testing.T) {
	s := newTestStore(t)

	want := []task.Task{testutil.MakeTask("1", priority.Medium, 25, "")}
	if err := s.WriteAllTasks(want); err != nil {
		t.Fatalf("WriteAllTasks: %v", err)
	}

	got, err := s.LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks: %v", err)
	}

	if len(got) != 1 || got[0].ID != want[0].ID {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestDeleteTask_RemovesMatchingID(t *testing.T) {
	tasks := []task.Task{
		testutil.MakeTask("1", priority.Medium, 25, ""),
		testutil.MakeTask("2", priority.Medium, 25, ""),
		testutil.MakeTask("3", priority.Medium, 25, ""),
	}

	got := task.Delete(tasks, "2")

	if len(got) != 2 {
		t.Fatalf("got %d tasks, want 2", len(got))
	}
	for _, tsk := range got {
		if tsk.ID == "2" {
			t.Error("task with ID 2 was not removed")
		}
	}
}

func TestDeleteTask_UnknownIDIsNoOp(t *testing.T) {
	tasks := []task.Task{testutil.MakeTask("1", priority.Medium, 25, "")}

	got := task.Delete(tasks, "does-not-exist")

	if len(got) != 1 {
		t.Errorf("got %d tasks, want 1 (unchanged)", len(got))
	}
}

func TestCompleteCurrentTask_MovesTaskToCompletedAndClearsCurrent(t *testing.T) {
	s := newTestStore(t)

	current := testutil.MakeTask("1", priority.High, 25, "work")
	if err := s.WriteCurrentTask(current); err != nil {
		t.Fatalf("WriteCurrentTask: %v", err)
	}

	if err := s.CompleteCurrentTask(); err != nil {
		t.Fatalf("CompleteCurrentTask: %v", err)
	}

	if _, err := os.Stat(s.CurrentTaskFilePath); !os.IsNotExist(err) {
		t.Errorf("expected current task file to be removed, stat err = %v", err)
	}

	completed, err := readTasksFile(s.CompletedTasksFilePath)
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
