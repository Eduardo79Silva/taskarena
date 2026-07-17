package tasklist

import (
	"strings"
)

func (m Model) View() string {
	if len(m.tasks) == 0 {
		return m.styles.Empty.Render("No tasks found.")
	}

	rows := make([]string, 0, len(m.tasks))

	for i, task := range m.tasks {
		row := NewRow(
			task,
			i == m.selected,
			m.width,
			m.styles,
		)

		rows = append(rows, row.View())
	}

	return strings.Join(rows, "\n")
}
