package cli

import (
	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func newDoneCmd(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Complete the current task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Done()
		},
	}
}
