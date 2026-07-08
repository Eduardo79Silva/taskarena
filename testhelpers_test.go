package main

import (
	"testing"
	"time"
)

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
