package cli

import (
	"fmt"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/spf13/cobra"
)

func newPushCmd(a *app.App) *cobra.Command {
	var p priority.Level
	var name, description, tag string
	var timeEstimate int

	cmd := &cobra.Command{
		Use:   "push",
		Short: "Add a new task",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := a.Push(name, description, timeEstimate, p, tag)
			if err != nil {
				return err
			}
			fmt.Println("added:", t.Name)
			return nil
		},
	}

	p = a.DefaultPriority()
	cmd.Flags().VarP(&p, "priority", "p", "priority (low|medium|high|veryhigh)")
	cmd.Flags().StringVarP(&name, "name", "n", "", "name")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "description")
	cmd.Flags().IntVarP(&timeEstimate, "time", "t", a.DefaultTimeEstimate(), "time estimate")
	cmd.Flags().StringVar(&tag, "tag", "", "tag")

	return cmd
}
