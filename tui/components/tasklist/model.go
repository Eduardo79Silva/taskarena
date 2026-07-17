package tasklist

import (
	"github.com/eduardo79silva/taskarena/internal/task"
	"github.com/eduardo79silva/taskarena/tui/styles"
)

type Model struct {
	tasks    []task.Task
	selected int
	width    int

	styles styles.Styles
}

func New(tasks []task.Task, styles styles.Styles) Model {
	return Model{
		tasks:  tasks,
		styles: styles,
	}
}

func (m *Model) SetTasks(tasks []task.Task) {
	m.tasks = tasks

	if m.selected >= len(tasks) {
		m.selected = len(tasks) - 1
	}
}

func (m *Model) MoveUp() {
	if m.selected > 0 {
		m.selected--
	}
}

func (m *Model) MoveDown() {
	if m.selected < len(m.tasks)-1 {
		m.selected++
	}
}

func (m Model) SelectedTask() (task.Task, bool) {
	if len(m.tasks) == 0 {
		return task.Task{}, false
	}

	return m.tasks[m.selected], true
}

func (m *Model) SetWidth(width int) {
	m.width = width
}
