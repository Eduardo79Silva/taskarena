package task

import "time"

func (t *Task) UpdateTime() {
	t.UpdatedAt = time.Now()
}

func (t *Task) CalculateTimeSpent() {
	elapsed := time.Since(t.UpdatedAt)
	if t.TimeSpent == nil {
		t.TimeSpent = new(time.Duration)
	}
	*t.TimeSpent += elapsed
}
