package tasks

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

func (m Model) View() tea.View {
	title := styles.Title.Render("Which task should we pick?\n\n")

	if len(m.tasks) == 0 {
		empty := styles.Normal.Render("No tasks found.\n")
		return tea.NewView(empty)
	}

	var rows []string

	for i, t := range m.tasks {
		if i == m.selected {
			rows = append(rows,
				styles.Selected.Render(t.Name),
			)
		} else {
			rows = append(rows,
				styles.Normal.Render(t.Name),
			)
		}
	}

	footer := styles.Title.Render("q quit • ? help")

	return tea.NewView(lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		strings.Join(rows, "\n"),
		footer,
	))
}
