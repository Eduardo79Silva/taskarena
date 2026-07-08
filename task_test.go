package main

import (
	"testing"
	"time"
)

func TestPriorityLevel_String(t *testing.T) {
	tests := []struct {
		name string
		p    PriorityLevel
		want string
	}{
		{"low", LowPriority, "low"},
		{"medium", MediumPriority, "medium"},
		{"high", HighPriority, "high"},
		{"veryhigh", VeryHighPriority, "veryhigh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.String()
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPriorityLevel_Set(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    PriorityLevel
		wantErr bool
	}{
		{"low", "low", LowPriority, false},
		{"medium", "medium", MediumPriority, false},
		{"high", "high", HighPriority, false},
		{"veryhigh", "veryhigh", VeryHighPriority, false},
		{"case insensitive", "HIGH", HighPriority, false},
		{"invalid", "urgent", LowPriority, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p PriorityLevel
			err := p.Set(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Set(%q) expected an error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("Set(%q) unexpected error: %v", tt.input, err)
			}
			if p != tt.want {
				t.Errorf("Set(%q) = %v, want %v", tt.input, p, tt.want)
			}
		})
	}
}

func TestCreateTask_EmptyNameReturnsError(t *testing.T) {
	_, err := createTask("", "desc", 25, MediumPriority, "work")
	if err == nil {
		t.Fatal("createTask with empty name: expected an error, got nil")
	}
}

func TestCreateTask_ValidInputPopulatesFields(t *testing.T) {
	before := time.Now()
	task, err := createTask("Write tests", "cover the core logic", 30, HighPriority, "work")
	after := time.Now()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID == "" {
		t.Error("expected a generated ID, got empty string")
	}
	if task.Name != "Write tests" {
		t.Errorf("Name = %q, want %q", task.Name, "Write tests")
	}
	if task.Description != "cover the core logic" {
		t.Errorf("Description = %q, want %q", task.Description, "cover the core logic")
	}
	if task.TimeEstimate != 30 {
		t.Errorf("TimeEstimate = %d, want 30", task.TimeEstimate)
	}
	if task.Priority != HighPriority {
		t.Errorf("Priority = %v, want %v", task.Priority, HighPriority)
	}
	if task.Tag != "work" {
		t.Errorf("Tag = %q, want %q", task.Tag, "work")
	}
	if task.CompletedAt != nil {
		t.Errorf("CompletedAt = %v, want nil", task.CompletedAt)
	}

	if task.CreatedAt.Before(before) || task.CreatedAt.After(after) {
		t.Errorf("CreatedAt = %v, want between %v and %v", task.CreatedAt, before, after)
	}
}

func TestCreateTask_GeneratesUniqueIDs(t *testing.T) {
	t1, err := createTask("a", "", 10, LowPriority, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t2, err := createTask("b", "", 10, LowPriority, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if t1.ID == t2.ID {
		t.Errorf("expected unique IDs, both were %q", t1.ID)
	}
}
