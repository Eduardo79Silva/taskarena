package app

import (
	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
)

func (a *App) Push(name, description string, timeEstimate int, p priority.Level, tag string) (task.Task, error) {
	t, err := task.New(name, description, timeEstimate, p, tag)
	if err != nil {
		return task.Task{}, err
	}
	if err := a.store.PushTask(t); err != nil {
		return task.Task{}, err
	}
	return t, nil
}

func (a *App) Pull(tagFilter string, timeFilter int) (task.Task, error) {
	tasks, err := a.store.LoadTasks()
	if err != nil {
		return task.Task{}, err
	}
	if tagFilter != "" {
		tasks = task.FilterByTag(tasks, tagFilter)
	}
	if timeFilter != -1 {
		tasks = task.FilterByTime(tasks, timeFilter)
	}

	newTask, err := a.scheduler.SelectNext(tasks)
	if err != nil {
		return task.Task{}, err
	}
	newTask.UpdateTime()

	var currentTask *task.Task
	currentTask, err = a.store.ReadCurrentTask()

	if currentTask != nil {
		currentTask.UpdateTime()
		currentTask.CalculateTimeSpent()
		if err := a.store.PushTask(*currentTask); err != nil {
			return task.Task{}, err
		}
		tasks, err = a.store.LoadTasks()
		if err != nil {
			return task.Task{}, err
		}
	}

	if err := a.store.WriteCurrentTask(newTask); err != nil {
		return task.Task{}, err
	}
	tasks = task.Delete(tasks, newTask.ID)
	if err := a.store.WriteAllTasks(tasks); err != nil {
		return task.Task{}, err
	}

	return newTask, nil
}

func (a *App) FinishCurrentTask() error {
	return a.store.CompleteCurrentTask()
}

func (a *App) FinishTask(task task.Task) error {
	return a.store.CompleteCurrentTask()
}
