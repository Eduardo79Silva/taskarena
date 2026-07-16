package cli

import (
	"fmt"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/internal/status"
	"github.com/spf13/cobra"
)

func newStatusCmd(a *app.App) *cobra.Command {
	var format string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show current task status",
		RunE: func(cmd *cobra.Command, args []string) error {
			view, err := a.CurrentTaskView()
			if err != nil {
				return err
			}

			switch format {
			case "waybar":
				out, err := status.FormatWaybar(view)
				if err != nil {
					return err
				}
				fmt.Println(out)
			case "notify":
				fmt.Println(status.FormatNotification(view))
			case "plain":
				fmt.Println(status.FormatPlain(view))
			default:
				return fmt.Errorf("unknown format: %s", format)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&format, "format", "plain", "output format (plain|waybar|notify)")
	return cmd
}
