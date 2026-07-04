package main

import (
	"testing"
	"time"
)

// makeTask builds a Task with sensible defaults, letting each test override
// only the fields it cares about. This keeps table-driven tests readable:
// you can write makeTask("t1", HighPriority, 25, "work") instead of a
// six-field struct literal every time.
func makeTask(id string, priority PriorityLevel, timeEstimate int, tag string) Task {
	return Task{
		ID:           id,
		Name:         "task-" + id,
		Description:  "description for " + id,
		TimeEstimate: timeEstimate,
		Priority:     priority,
		Tag:          tag,
		CreatedAt:    time.Now(),
	}
}

// withTempStoragePaths points the package's file-path variables at a
// temporary directory for the duration of a test, then restores the
// originals. storage.go keeps these as package-level vars (rather than
// consts) specifically so tests can swap them out like this.
//
// Not safe to use with t.Parallel(): it mutates shared package state.
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
