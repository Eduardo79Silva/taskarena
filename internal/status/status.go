package status

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/eduardo79silva/taskarena/internal/store"
	"github.com/eduardo79silva/taskarena/internal/task"
)

type WaybarOutput struct {
	Text       string `json:"text"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

func getCurrentTaskView() (*task.CurrentTaskView, error) {
	currentTask, err := store.ReadTaskFile(store.CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	view := task.CurrentTaskView{}

	view.Name = currentTask.Name
	view.Description = currentTask.Description
	view.Priority = currentTask.Priority
	view.TimeEstimate = currentTask.TimeEstimate

	return &view, nil
}

func formatWaybar(view *task.CurrentTaskView) (string, error) {
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

func formatNotification(view *task.CurrentTaskView) string {
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

func formatPlain(view *task.CurrentTaskView) string {
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
