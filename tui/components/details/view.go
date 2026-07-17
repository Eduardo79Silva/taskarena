package details

import (
	"charm.land/lipgloss/v2"
	"github.com/eduardo79silva/taskarena/internal/task"
)

func field(m Model, label, value string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.styles.Label.Render(label+":"),
		" ",
		m.styles.Value.Render(value),
	)
}

func (m Model) View() string {
	t := m.presentedTask

	if (t == task.Task{}) {
		return m.styles.Empty.Render("No task selected.")
	}

	duration := "-"
	if t.TimeSpent != nil {
		duration = t.TimeSpent.String()
	}

	completed := "-"
	if t.CompletedAt != nil {
		completed = t.CompletedAt.String()
	}

	description := "-"
	if t.Description != "" {
		description = t.Description
	}

	tag := "-"
	if t.Tag != "" {
		tag = t.Tag
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.Title.Render(t.Name),
		"",
		field(m, "Priority", t.Priority.String()),
		field(m, "Duration", duration),
		field(m, "Completed", completed),
		field(m, "Tag", tag),
		"",
		m.styles.Normal.Render(description),
	)
}
