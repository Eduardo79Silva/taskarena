package tasks

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// The header
	var s strings.Builder
	s.WriteString("Which task should we pick?\n\n")

	if len(m.tasks) == 0 {
		s.WriteString("No tasks found.\n")
		return tea.NewView(s.String())
	}

	// Iterate over our choices
	for i, t := range m.tasks {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.selected == i {
			cursor = ">" // cursor!
		}

		// Render the row
		fmt.Fprintf(&s, "%s [%s] %s\n", cursor, t.Tag, t.Name)
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}
