package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type PriorityLevel int

const (
	LowPriority PriorityLevel = iota
	MediumPriority
	HighPriority
	VeryHighPriority
)

func (p *PriorityLevel) String() string {
	return [...]string{"low", "medium", "high", "veryhigh"}[*p]
}

func (p *PriorityLevel) Set(s string) error {
	switch strings.ToLower(s) {
	case "low":
		*p = LowPriority
	case "medium":
		*p = MediumPriority
	case "high":
		*p = HighPriority
	case "veryhigh":
		*p = VeryHighPriority
	default:
		return fmt.Errorf("invalid priority: %s", s)
	}
	return nil
}

type Task struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	TimeEstimate int           `json:"timeEstimate"`
	Priority     PriorityLevel `json:"priority"`
	CreatedAt    time.Time     `json:"created_at"`
	CompletedAt  *time.Time    `json:"completed_at,omitempty"`
}

func createTask(name string, description string, timeEstimate int, priority PriorityLevel) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}
	return Task{uuid.New().String(), name, description, timeEstimate, priority, time.Now(), nil}, nil
}
