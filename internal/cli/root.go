package cli

import (
	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCmd(a *app.App) *cobra.Command {
	root := &cobra.Command{
		Use:   "taskarena",
		Short: "A weighted-scheduling task manager",
	}

	root.AddCommand(
		newPushCmd(a),
		newPullCmd(a),
		newDoneCmd(a),
		newStatusCmd(a),
		newListCmd(a),
		newEditCmd(a),
	)

	return root
}
