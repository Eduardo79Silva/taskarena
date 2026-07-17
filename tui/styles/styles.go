package styles

import "charm.land/lipgloss/v2"

type Styles struct {
	Title    lipgloss.Style
	Selected lipgloss.Style
	Normal   lipgloss.Style
	Empty    lipgloss.Style
	Footer   lipgloss.Style
}

func DefaultStyles() Styles {
	return Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86")),

		Selected: lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")),

		Normal: lipgloss.NewStyle(),
		Empty: lipgloss.NewStyle().
			Faint(true),

		Footer: lipgloss.NewStyle().
			Faint(true),
	}
}
