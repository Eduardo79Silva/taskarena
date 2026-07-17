package tasklist

import (
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

type Row struct {
	task     task.Task
	selected bool
	width    int
	styles   styles.Styles
}

func NewRow(
	task task.Task,
	selected bool,
	width int,
	styles styles.Styles,
) Row {
	return Row{
		task:     task,
		selected: selected,
		width:    width,
		styles:   styles,
	}
}

func (r Row) View() string {
	content := "  " + r.task.Name

	style := r.styles.Normal

	if r.selected {
		style = r.styles.Selected
		content = "> " + r.task.Name
	}

	return style.
		Width(r.width).
		Render(content)
}
