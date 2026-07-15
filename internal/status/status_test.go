package status

import (
	"encoding/json"
	"testing"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
)

func TestFormatWaybar_NilView(t *testing.T) {
	out, err := FormatWaybar(nil)
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
	view := &task.CurrentTaskView{Name: "Write tests", Description: "cover the core logic"}
	out, err := FormatWaybar(view)
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
		fn   func(*task.CurrentTaskView) string
	}{
		{"FormatNotification", FormatNotification},
		{"FormatPlain", FormatPlain},
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
	view := &task.CurrentTaskView{
		Name:         "Write tests",
		Description:  "cover the core logic",
		Priority:     priority.High,
		TimeEstimate: 30,
	}
	got := FormatNotification(view)
	want := "Write tests (high, 30m)\ncover the core logic"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormatPlain_WithView(t *testing.T) {
	view := &task.CurrentTaskView{
		Name:         "Write tests",
		Priority:     priority.High,
		TimeEstimate: 30,
	}
	got := FormatPlain(view)
	want := "Write tests (high, 30m)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
