package cli

import (
	"fmt"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func newPullCmd(a *app.App) *cobra.Command {
	var tagFilter string
	var timeFilter int

	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull the next task",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := a.Pull(tagFilter, timeFilter)
			if err != nil {
				return err
			}
			fmt.Println(t)
			return nil
		},
	}

	cmd.Flags().StringVar(&tagFilter, "tag", "", "filter by tag")
	cmd.Flags().IntVar(&timeFilter, "time", -1, "filter by time")

	return cmd
}
