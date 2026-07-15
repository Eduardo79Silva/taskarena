package testutil

import (
	"time"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
)

func MakeTask(id string, priority priority.Level, timeEstimate int, tag string) task.Task {
	return task.Task{
		ID:           id,
		Name:         "task-" + id,
		Description:  "description for " + id,
		TimeEstimate: timeEstimate,
		Priority:     priority,
		Tag:          tag,
		CreatedAt:    time.Now(),
	}
}
