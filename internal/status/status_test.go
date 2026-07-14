package main

import (
	"encoding/json"
	"testing"
)

func TestGetCurrentTaskView_NoCurrentTaskReturnsNil(t *testing.T) {
	withTempStoragePaths(t)

	view, err := getCurrentTaskView()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if view != nil {
		t.Errorf("got %+v, want nil", view)
	}
}

func TestGetCurrentTaskView_ReturnsCurrentTaskFields(t *testing.T) {
	withTempStoragePaths(t)

	task := makeTask("1", HighPriority, 30, "work")
	writeCurrentTask(task)

	view, err := getCurrentTaskView()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if view == nil {
		t.Fatal("got nil view, want a populated one")
	}
	if view.Name != task.Name || view.Priority != task.Priority || view.TimeEstimate != task.TimeEstimate {
		t.Errorf("view = %+v, want fields matching task %+v", view, task)
	}
}

func TestFormatWaybar_NilView(t *testing.T) {
	out, err := formatWaybar(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed WaybarOutput
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if parsed.Text != "No Task" {
		t.Errorf("Text = %q, want %q", parsed.Text, "No Task")
	}
	if parsed.Class != "idle" {
		t.Errorf("Class = %q, want %q", parsed.Class, "idle")
	}
}

func TestFormatWaybar_WithView(t *testing.T) {
	view := &CurrentTaskView{Name: "Write tests", Description: "cover the core logic"}

	out, err := formatWaybar(view)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed WaybarOutput
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if parsed.Text != view.Name {
		t.Errorf("Text = %q, want %q", parsed.Text, view.Name)
	}
	if parsed.Tooltip != view.Description {
		t.Errorf("Tooltip = %q, want %q", parsed.Tooltip, view.Description)
	}
	if parsed.Class != "task" {
		t.Errorf("Class = %q, want %q", parsed.Class, "task")
	}
}

func TestFormatters_NilView(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*CurrentTaskView) string
	}{
		{"formatNotification", formatNotification},
		{"formatPlain", formatPlain},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(nil)
			want := "There isn't any active task currently"
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestFormatNotification_WithView(t *testing.T) {
	view := &CurrentTaskView{
		Name:         "Write tests",
		Description:  "cover the core logic",
		Priority:     HighPriority,
		TimeEstimate: 30,
	}

	got := formatNotification(view)
	want := "Write tests (high, 30m)\ncover the core logic"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormatPlain_WithView(t *testing.T) {
	view := &CurrentTaskView{
		Name:         "Write tests",
		Priority:     HighPriority,
		TimeEstimate: 30,
	}

	got := formatPlain(view)
	want := "Write tests (high, 30m)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
