package main

import (
	"testing"
	"time"
)

// --- filterTasksByTag ---
func TestFilterTasksByTag(t *testing.T) {
	tasks := []Task{
		makeTask("1", MediumPriority, 25, "work"),
		makeTask("2", MediumPriority, 25, "home"),
		makeTask("3", MediumPriority, 25, "work"),
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

func TestFilterTasksByTag_NoMatches(t *testing.T) {
	tasks := []Task{makeTask("1", MediumPriority, 25, "work")}

	got := filterTasksByTag(tasks, "home")

	if len(got) != 0 {
		t.Errorf("got %d tasks, want 0", len(got))
	}
}

// --- wsmScore ---
//
// wsmScore blends three normalized factors (priority, time, age). Rather
// than pin down an exact float value (brittle, and not what we actually
// care about), these tests check the *direction* each factor pushes the
// score: that's the actual behavior the scheduler depends on.
func TestWsmScore_HigherPriorityScoresHigher(t *testing.T) {
	low := makeTask("low", LowPriority, 25, "")
	high := makeTask("high", HighPriority, 25, "")
	// Same CreatedAt so age contributes equally to both.
	now := time.Now()
	low.CreatedAt = now
	high.CreatedAt = now

	lowScore := wsmScore(low, 25, 25)
	highScore := wsmScore(high, 25, 25)

	if highScore <= lowScore {
		t.Errorf("high priority score (%v) should be greater than low priority score (%v)", highScore, lowScore)
	}
}

func TestWsmScore_ShorterTaskScoresHigher(t *testing.T) {
	now := time.Now()
	short := makeTask("short", MediumPriority, 10, "")
	long := makeTask("long", MediumPriority, 50, "")
	short.CreatedAt = now
	long.CreatedAt = now

	minTime, maxTime := 10, 50

	shortScore := wsmScore(short, minTime, maxTime)
	longScore := wsmScore(long, minTime, maxTime)

	if shortScore <= longScore {
		t.Errorf("shorter task score (%v) should be greater than longer task score (%v)", shortScore, longScore)
	}
}

func TestWsmScore_OlderTaskScoresHigher(t *testing.T) {
	older := makeTask("older", MediumPriority, 25, "")
	newer := makeTask("newer", MediumPriority, 25, "")
	older.CreatedAt = time.Now().Add(-48 * time.Hour)
	newer.CreatedAt = time.Now()

	olderScore := wsmScore(older, 25, 25)
	newerScore := wsmScore(newer, 25, 25)

	if olderScore <= newerScore {
		t.Errorf("older task score (%v) should be greater than newer task score (%v)", olderScore, newerScore)
	}
}

func TestWsmScore_SingleTimeValueDoesNotDivideByZero(t *testing.T) {
	// When every task has the same estimate, minTime == maxTime. The
	// function has an explicit guard for this; this test pins that
	// behavior down so it can't regress into a division by zero.
	task := makeTask("1", MediumPriority, 25, "")
	score := wsmScore(task, 25, 25)

	if score <= 0 {
		t.Errorf("expected a positive score when minTime == maxTime, got %v", score)
	}
}

// --- selectNextTask ---
func TestSelectNextTask_EmptyListReturnsError(t *testing.T) {
	_, err := selectNextTask(nil)
	if err != errEmptyTaskList {
		t.Errorf("got error %v, want %v", err, errEmptyTaskList)
	}
}

func TestSelectNextTask_SingleTaskIsAlwaysReturned(t *testing.T) {
	task := makeTask("only", MediumPriority, 25, "")

	got, err := selectNextTask([]Task{task})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != task.ID {
		t.Errorf("got task %q, want %q", got.ID, task.ID)
	}
}

// selectNextTask is randomized by design, so we can't assert which task
// comes back on a given run. What we *can* assert, many times over, is
// the invariant that matters: the result always comes from the input
// list. That's what callers actually rely on.
func TestSelectNextTask_AlwaysReturnsATaskFromTheList(t *testing.T) {
	tasks := []Task{
		makeTask("1", LowPriority, 10, ""),
		makeTask("2", MediumPriority, 25, ""),
		makeTask("3", HighPriority, 40, ""),
		makeTask("4", VeryHighPriority, 5, ""),
	}

	valid := make(map[string]bool, len(tasks))
	for _, task := range tasks {
		valid[task.ID] = true
	}

	for i := 0; i < 100; i++ {
		got, err := selectNextTask(tasks)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !valid[got.ID] {
			t.Fatalf("selectNextTask returned unknown task ID %q", got.ID)
		}
	}
}
