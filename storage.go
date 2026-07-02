package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var TasksFilePath = defaultTasksFilePath()
var Tasks []Task

func defaultTasksFilePath() string {
	home, err := os.UserHomeDir()
	check(err)
	dir := filepath.Join(home, ".config", "taskarena")
	err = os.MkdirAll(dir, 0755)
	check(err)
	return filepath.Join(dir, "tasks.json")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func appendTaskToJsonString(task Task, jsonText string) string {
	err := json.Unmarshal([]byte(jsonText), &Tasks)
	check(err)
	Tasks = append(Tasks, task)
	result, err := json.MarshalIndent(Tasks, "", "\t")
	check(err)
	return string(result)
}

func readFile(filePath string) string {
	dat, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return ""
	}
	check(err)
	return string(dat)
}

func writeTask(task Task) {
	jsonText := readFile(TasksFilePath)
	if jsonText == "" {
		jsonText = "[]"
	}
	parsedJson := appendTaskToJsonString(task, jsonText)
	err := os.WriteFile(TasksFilePath, []byte(parsedJson), 0644)
	check(err)
}

func writeAllTasks(tasks []Task) {
	result, err := json.MarshalIndent(tasks, "", "\t")
	check(err)

	parsedJson := string(result)
	err = os.WriteFile(TasksFilePath, []byte(parsedJson), 0644)
	check(err)
}

func deleteTask(tasks []*Task, taskID string) []*Task {
	for i, task := range tasks {
		if task.ID == taskID {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func loadTasks() []Task {
	jsonText := readFile(TasksFilePath)

	if jsonText == "" {
		return []Task{}
	}

	err := json.Unmarshal([]byte(jsonText), &Tasks)
	check(err)

	return Tasks
}
