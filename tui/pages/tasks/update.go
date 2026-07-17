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

		m.taskList.SetTasks(msg.Tasks)

	case TaskCompletedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

		return m, loadTasksCmd(m.app)

	case tea.WindowSizeMsg:
		m.taskList.SetWidth(msg.Width)

	case tea.KeyMsg:
		switch msg.String() {

		case "down", "j":
			m.taskList.MoveDown()

		case "up", "k":
			m.taskList.MoveUp()

		case "d":
			if _, ok := m.taskList.SelectedTask(); ok {
				return m, completeTaskCmd(m.app)
			}

		case "r":
			return m, loadTasksCmd(m.app)
		}
	}

	return m, nil
}
