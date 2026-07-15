package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func newListCmd(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := a.ListTasks()
			if err != nil {
				return err
			}
			if len(tasks) == 0 {
				fmt.Println("No tasks.")
				return nil
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
			fmt.Fprintln(w, "TAG\tPRIORITY\tTIME\tNAME")
			for _, t := range tasks {
				fmt.Fprintf(w, "%s\t%s\t%dm\t%s\n", t.Tag, t.Priority.String(), t.TimeEstimate, t.Name)
			}
			return w.Flush()
		},
	}
}
