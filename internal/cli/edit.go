package cli

import (
	"os"
	"os/exec"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/spf13/cobra"
)

func newEditCmd(a *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit the tasks file directly",
		RunE: func(cmd *cobra.Command, args []string) error {
			return editFile(a.TasksFilePath())
		},
	}
}

func editFile(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
