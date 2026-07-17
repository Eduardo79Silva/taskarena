package details

import (
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

type Model struct {
	presentedTask task.Task
	selected      int
	width         int

	styles styles.Styles
}

func New(t task.Task, styles styles.Styles) Model {
	return Model{
		presentedTask: t,
		styles:        styles,
	}
}

func (m *Model) SetTask(t task.Task) {
	m.presentedTask = t
}
