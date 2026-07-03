package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type WaybarOutput struct {
	Text       string `json:"text"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

func getCurrentTaskView() (*CurrentTaskView, error) {
	task, err := readTaskFile(CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	view := CurrentTaskView{}

	view.Name = task.Name
	view.Description = task.Description
	view.Priority = task.Priority
	view.TimeEstimate = task.TimeEstimate

	return &view, nil
}

func formatWaybar(view *CurrentTaskView) (string, error) {
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

func formatNotification(view *CurrentTaskView) string {
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

func formatPlain(view *CurrentTaskView) string {
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
