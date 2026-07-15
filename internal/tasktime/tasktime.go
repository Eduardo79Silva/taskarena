package tasktime

import (
	"github.com/eduardo79silva/taskarena/internal/task"
	"time"
)

func UpdateTaskTime(task *task.Task) {
	task.UpdatedAt = time.Now()
}

func CalculateTimeSpent(task *task.Task) {
	lastUpdate := task.UpdatedAt
	timeSpent := time.Since(lastUpdate)

	if task.TimeSpent == nil {
		task.TimeSpent = new(time.Duration)
	}

	*task.TimeSpent += timeSpent
}
