package cli

import (
	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCmd(a *app.App, version string) *cobra.Command {
	root := &cobra.Command{
		Use:     "taskarena",
		Short:   "A weighted-scheduling task manager",
		Version: version,
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
