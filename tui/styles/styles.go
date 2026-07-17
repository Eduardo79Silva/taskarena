package styles

import "charm.land/lipgloss/v2"

var Title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("205")).
	Padding(0, 1)

var Selected = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)

var Normal = lipgloss.NewStyle()
