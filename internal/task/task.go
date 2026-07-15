package task

import (
	"errors"
	"time"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/google/uuid"
)

type Task struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	TimeEstimate int            `json:"timeEstimate"`
	TimeSpent    *time.Duration `json:"timeSpent,omitempty"`
	Priority     priority.Level `json:"priority"`
	Tag          string         `json:"tag"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CompletedAt  *time.Time     `json:"completed_at,omitempty"`
}

type CurrentTaskView struct {
	Name         string
	Description  string
	Priority     priority.Level
	TimeEstimate int
}

func createTask(name string, description string, timeEstimate int, priority priority.Level, tag string) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}
	return Task{uuid.New().String(), name, description, timeEstimate, nil, priority, tag, time.Now(), time.Now(), nil}, nil
}
