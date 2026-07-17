package tasks

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

func (m Model) View() tea.View {
	title := styles.DefaultStyles().Title.Render("Which task should we pick?\n\n")

	footer := styles.DefaultStyles().Title.Render("q quit • ? help")

	return tea.NewView(lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		m.taskList.View(),
		footer,
	))
}
