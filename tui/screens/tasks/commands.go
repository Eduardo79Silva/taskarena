package tasks

import (
	tea "charm.land/bubbletea/v2"

	"github.com/eduardo79silva/taskarena/internal/app"
)

func loadTasksCmd(a *app.App) tea.Cmd {
	return func() tea.Msg {
		tasks, err := a.ListTasks()

		return TasksLoadedMsg{
			Tasks: tasks,
			Err:   err,
		}
	}
}

func completeTaskCmd(a *app.App) tea.Cmd {
	return func() tea.Msg {
		err := a.Done()

		return TaskCompletedMsg{
			Err: err,
		}
	}
}
