package app

import (
	"github.com/eduardo79silva/taskarena/internal/task"
)

func (a *App) CurrentTaskView() (*task.CurrentTaskView, error) {
	return a.store.GetCurrentTaskView()
}

func (a *App) ListTasks() ([]task.Task, error) {
	return a.store.LoadTasks()
}

func (a *App) TasksFilePath() string {
	return a.store.TasksFilePath
}
