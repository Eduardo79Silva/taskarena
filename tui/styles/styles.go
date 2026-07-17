package styles

import "charm.land/lipgloss/v2"

type Styles struct {
	Title    lipgloss.Style
	Label    lipgloss.Style
	Value    lipgloss.Style
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
		Label: lipgloss.NewStyle().
			Bold(true),
		Value: lipgloss.NewStyle(),

		Selected: lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			PaddingLeft(1).
			PaddingRight(1),

		Normal: lipgloss.NewStyle(),
		Empty: lipgloss.NewStyle().
			Faint(true),

		Footer: lipgloss.NewStyle().
			Faint(true),
	}
}
