package tasklist

import (
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

type Row struct {
	task     task.Task
	selected bool
	styles   styles.Styles
}

func NewRow(
	task task.Task,
	selected bool,
	styles styles.Styles,
) Row {
	return Row{
		task:     task,
		selected: selected,
		styles:   styles,
	}
}

func (r Row) View() string {
	title := r.task.Name

	if r.selected {
		return r.styles.Selected.Render("> " + title)
	}

	return r.styles.Normal.Render("  " + title)
}
