package cli

import (
	tea "charm.land/bubbletea/v2"
	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/tui"
	"github.com/spf13/cobra"
)

func newTUICommand(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use: "tui",
		RunE: func(cmd *cobra.Command, args []string) error {

			p := tea.NewProgram(tui.New(a))

			_, err := p.Run()

			return err
		},
	}
}
