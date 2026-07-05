package main

import "time"

func updateTaskTime(task *Task) {
	task.UpdatedAt = time.Now()
}

func calculateTimeSpent(task *Task) {
	lastUpdate := task.UpdatedAt
	timeSpent := time.Since(lastUpdate)

	if task.TimeSpent == nil {
		task.TimeSpent = new(time.Duration)
	}

	*task.TimeSpent += timeSpent
}
