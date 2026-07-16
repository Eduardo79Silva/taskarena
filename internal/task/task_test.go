package task

import (
	"testing"
	"time"

	"github.com/eduardo79silva/taskarena/internal/priority"
)

func TestPriorityLevel_String(t *testing.T) {
	tests := []struct {
		name string
		p    priority.Level
		want string
	}{
		{"low", priority.Low, "low"},
		{"medium", priority.Medium, "medium"},
		{"high", priority.High, "high"},
		{"veryhigh", priority.VeryHigh, "veryhigh"},
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
		want    priority.Level
		wantErr bool
	}{
		{"low", "low", priority.Low, false},
		{"medium", "medium", priority.Medium, false},
		{"high", "high", priority.High, false},
		{"veryhigh", "veryhigh", priority.VeryHigh, false},
		{"case insensitive", "HIGH", priority.High, false},
		{"invalid", "urgent", priority.Low, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p priority.Level
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
	_, err := New("", "desc", 25, priority.Medium, "work")
	if err == nil {
		t.Fatal("New with empty name: expected an error, got nil")
	}
}

func TestCreateTask_ValidInputPopulatesFields(t *testing.T) {
	before := time.Now()
	tsk, err := New("Write tests", "cover the core logic", 30, priority.High, "work")
	after := time.Now()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tsk.ID == "" {
		t.Error("expected a generated ID, got empty string")
	}
	if tsk.Name != "Write tests" {
		t.Errorf("Name = %q, want %q", tsk.Name, "Write tests")
	}
	if tsk.Description != "cover the core logic" {
		t.Errorf("Description = %q, want %q", tsk.Description, "cover the core logic")
	}
	if tsk.TimeEstimate != 30 {
		t.Errorf("TimeEstimate = %d, want 30", tsk.TimeEstimate)
	}
	if tsk.Priority != priority.High {
		t.Errorf("Priority = %v, want %v", tsk.Priority, priority.High)
	}
	if tsk.Tag != "work" {
		t.Errorf("Tag = %q, want %q", tsk.Tag, "work")
	}
	if tsk.CompletedAt != nil {
		t.Errorf("CompletedAt = %v, want nil", tsk.CompletedAt)
	}
	if tsk.CreatedAt.Before(before) || tsk.CreatedAt.After(after) {
		t.Errorf("CreatedAt = %v, want between %v and %v", tsk.CreatedAt, before, after)
	}
}

func TestCreateTask_GeneratesUniqueIDs(t *testing.T) {
	t1, err := New("a", "", 10, priority.Low, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t2, err := New("b", "", 10, priority.Low, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if t1.ID == t2.ID {
		t.Errorf("expected unique IDs, both were %q", t1.ID)
	}
}
