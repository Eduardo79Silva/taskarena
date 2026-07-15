package status

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/eduardo79silva/taskarena/internal/task"
)

type WaybarOutput struct {
	Text       string `json:"text"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

func FormatWaybar(view *task.CurrentTaskView) (string, error) {
	out := WaybarOutput{}
	if view == nil {
		out.Text = "No Task"
		out.Tooltip = "There isn't any active task currently"
		out.Class = "idle"
	} else {
		out.Text = view.Name
		out.Tooltip = view.Description
		out.Class = "task"
	}

	data, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FormatNotification(view *task.CurrentTaskView) string {
	if view == nil {
		return "There isn't any active task currently"
	}
	builder := strings.Builder{}
	builder.WriteString(view.Name)
	builder.WriteString(" (")
	builder.WriteString(view.Priority.String())
	builder.WriteString(", ")
	builder.WriteString(strconv.Itoa(view.TimeEstimate))
	builder.WriteString("m)\n")
	builder.WriteString(view.Description)
	return builder.String()
}

func FormatPlain(view *task.CurrentTaskView) string {
	if view == nil {
		return "There isn't any active task currently"
	}
	builder := strings.Builder{}
	builder.WriteString(view.Name)
	builder.WriteString(" (")
	builder.WriteString(view.Priority.String())
	builder.WriteString(", ")
	builder.WriteString(strconv.Itoa(view.TimeEstimate))
	builder.WriteString("m)")
	return builder.String()
}
