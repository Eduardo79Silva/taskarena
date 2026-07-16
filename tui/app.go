package tui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/tui/screens/tasks"
)

type Model struct {
	app *app.App

	current Screen
}

func New(a *app.App) Model {
	tasks := tasks.New(a)

	return Model{
		app:     a,
		current: tasks,
	}
}

func (m Model) Init() tea.Cmd {
	return m.current.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}

	var cmd tea.Cmd
	m.current, cmd = m.current.Update(msg)

	return m, cmd
}

func (m Model) View() tea.View {
	return m.current.View()
}
