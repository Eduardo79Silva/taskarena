package tasks

import (
	tea "charm.land/bubbletea/v2"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/tui/components/tasklist"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

type Model struct {
	app *app.App

	taskList tasklist.Model

	err error
}

func New(app *app.App) Model {
	return Model{
		app:      app,
		taskList: tasklist.New(nil, styles.DefaultStyles()),
	}
}

func (m Model) Init() tea.Cmd {
	return loadTasksCmd(m.app)
}
