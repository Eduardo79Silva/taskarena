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

func New(name, description string, timeEstimate int, p priority.Level, tag string) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}
	now := time.Now()
	return Task{
		ID:           uuid.New().String(),
		Name:         name,
		Description:  description,
		TimeEstimate: timeEstimate,
		Priority:     p,
		Tag:          tag,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

func FilterByTag(tasks []Task, tag string) []Task {
	var filtered []Task
	for _, t := range tasks {
		if t.Tag == tag {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func FilterByTime(tasks []Task, timeLimit int) []Task {
	var filtered []Task
	for _, t := range tasks {
		if t.TimeEstimate <= timeLimit {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func Delete(tasks []Task, id string) []Task {
	for i, t := range tasks {
		if t.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}
