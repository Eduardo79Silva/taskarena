package tasks

import "github.com/eduardo79silva/taskarena/internal/task"

type TasksLoadedMsg struct {
	Tasks []task.Task
	Err   error
}

type TaskCompletedMsg struct {
	Err error
}
