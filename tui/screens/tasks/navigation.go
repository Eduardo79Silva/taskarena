package tasks

func (m *Model) moveDown() {
	if m.selected < len(m.tasks)-1 {
		m.selected++
	}
}

func (m *Model) moveUp() {
	if m.selected > 0 {
		m.selected--
	}
}
