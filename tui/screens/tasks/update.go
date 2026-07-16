package tasks

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case TasksLoadedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		m.tasks = msg.Tasks

	case TaskCompletedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		return m, loadTasksCmd(m.app)

	case tea.KeyMsg:
		switch msg.String() {

		case "down", "j":
			m.moveDown()

		case "up", "k":
			m.moveUp()

		case "d":
			if _, ok := m.SelectedTask(); ok {
				return m, completeTaskCmd(m.app)
			}

		case "r":
			return m, loadTasksCmd(m.app)
		}
	}

	return m, nil
}
