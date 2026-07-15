package scheduler

import (
	"testing"
	"time"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/internal/testutil"
)

func TestFilterTasksByTag(t *testing.T) {
	tasks := []task.Task{
		testutil.MakeTask("1", priority.Medium, 25, "work"),
		testutil.MakeTask("2", priority.Medium, 25, "home"),
		testutil.MakeTask("3", priority.Medium, 25, "work"),
	}

	got := filterTasksByTag(tasks, "work")

	if len(got) != 2 {
		t.Fatalf("got %d tasks, want 2", len(got))
	}
	for _, task := range got {
		if task.Tag != "work" {
			t.Errorf("filterTasksByTag returned task with tag %q, want %q", task.Tag, "work")
		}
	}
}

func TestFilterTasksByTime(t *testing.T) {
	tasks := []task.Task{
		testutil.MakeTask("1", priority.Medium, 50, "work"),
		testutil.MakeTask("2", priority.Medium, 25, "home"),
		testutil.MakeTask("3", priority.Medium, 5, "work"),
	}

	got := filterTasksByTime(tasks, 25)

	if len(got) != 2 {
		t.Fatalf("got %d tasks, want 2", len(got))
	}

	for _, task := range got {
		if task.TimeEstimate > 25 {
			t.Errorf("filterTasksByTag returned task with tag %q, want %q", task.Tag, "work")
		}
	}
}

func TestFilterTasksByTag_NoMatches(t *testing.T) {
	tasks := []task.Task{testutil.MakeTask("1", priority.Medium, 25, "work")}

	got := filterTasksByTag(tasks, "home")

	if len(got) != 0 {
		t.Errorf("got %d tasks, want 0", len(got))
	}
}

func TestWsmScore_HigherPriorityScoresHigher(t *testing.T) {
	low := testutil.MakeTask("low", priority.Low, 25, "")
	high := testutil.MakeTask("high", priority.High, 25, "")

	now := time.Now()
	low.CreatedAt = now
	high.CreatedAt = now

	s := New(testutil.MakeSchedulerConfig())

	lowScore := s.wsmScore(low, 25, 25)
	highScore := s.wsmScore(high, 25, 25)

	if highScore <= lowScore {
		t.Errorf("high priority score (%v) should be greater than low priority score (%v)", highScore, lowScore)
	}
}

func TestWsmScore_ShorterTaskScoresHigher(t *testing.T) {
	now := time.Now()
	short := testutil.MakeTask("short", priority.Medium, 10, "")
	long := testutil.MakeTask("long", priority.Medium, 50, "")
	short.CreatedAt = now
	long.CreatedAt = now

	minTime, maxTime := 10, 50

	s := New(testutil.MakeSchedulerConfig())

	shortScore := s.wsmScore(short, minTime, maxTime)
	longScore := s.wsmScore(long, minTime, maxTime)

	if shortScore <= longScore {
		t.Errorf("shorter task score (%v) should be greater than longer task score (%v)", shortScore, longScore)
	}
}

func TestWsmScore_OlderTaskScoresHigher(t *testing.T) {
	older := testutil.MakeTask("older", priority.Medium, 25, "")
	newer := testutil.MakeTask("newer", priority.Medium, 25, "")
	older.CreatedAt = time.Now().Add(-48 * time.Hour)
	newer.CreatedAt = time.Now()

	s := New(testutil.MakeSchedulerConfig())

	olderScore := s.wsmScore(older, 25, 25)
	newerScore := s.wsmScore(newer, 25, 25)

	if olderScore <= newerScore {
		t.Errorf("older task score (%v) should be greater than newer task score (%v)", olderScore, newerScore)
	}
}

func TestWsmScore_SingleTimeValueDoesNotDivideByZero(t *testing.T) {

	s := New(testutil.MakeSchedulerConfig())

	task := testutil.MakeTask("1", priority.Medium, 25, "")
	score := s.wsmScore(task, 25, 25)

	if score <= 0 {
		t.Errorf("expected a positive score when minTime == maxTime, got %v", score)
	}
}

func TestSelectNextTask_EmptyListReturnsError(t *testing.T) {
	s := New(testutil.MakeSchedulerConfig())

	_, err := s.SelectNextTask(nil)
	if err != errEmptyTaskList {
		t.Errorf("got error %v, want %v", err, errEmptyTaskList)
	}
}

func TestSelectNextTask_SingleTaskIsAlwaysReturned(t *testing.T) {
	s := New(testutil.MakeSchedulerConfig())

	testTask := testutil.MakeTask("only", priority.Medium, 25, "")

	got, err := s.SelectNextTask([]task.Task{testTask})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != testTask.ID {
		t.Errorf("got task %q, want %q", got.ID, testTask.ID)
	}
}

func TestSelectNextTask_AlwaysReturnsATaskFromTheList(t *testing.T) {
	tasks := []task.Task{
		testutil.MakeTask("1", priority.Low, 10, ""),
		testutil.MakeTask("2", priority.Medium, 25, ""),
		testutil.MakeTask("3", priority.High, 40, ""),
		testutil.MakeTask("4", priority.VeryHigh, 5, ""),
	}

	valid := make(map[string]bool, len(tasks))
	for _, task := range tasks {
		valid[task.ID] = true
	}

	s := New(testutil.MakeSchedulerConfig())

	for range 100 {
		got, err := s.SelectNextTask(tasks)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !valid[got.ID] {
			t.Fatalf("selectNextTask returned unknown task ID %q", got.ID)
		}
	}
}
