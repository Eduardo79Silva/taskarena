package tasks

import (
	tea "charm.land/bubbletea/v2"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/internal/task"
)

type Model struct {
	app *app.App

	tasks []task.Task

	selected int

	err error
}

func New(app *app.App) Model {
	return Model{
		app: app,
	}
}

func (m Model) Init() tea.Cmd {
	return loadTasksCmd(m.app)
}

func (m *Model) SelectedTask() (task.Task, bool) {
	if len(m.tasks) == 0 {
		return task.Task{}, false
	}

	return m.tasks[m.selected], true
}
