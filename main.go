package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type PriorityLevel int

const (
	LowPriority PriorityLevel = iota
	MediumPriority
	HighPriority
	VeryHighPriority
)

type Task struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	TimeEstimate int           `json:"timeEstimate"`
	Priority     PriorityLevel `json:"priority"`
}

var TasksFilePath = "./tasks.json"

var Tasks []Task

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func appendTaskToJsonString(task Task, jsonText string) string {
	fmt.Println(jsonText)

	err := json.Unmarshal([]byte(jsonText), &Tasks)

	check(err)

	Tasks = append(Tasks, task)

	result, err := json.MarshalIndent(Tasks, "", "\t")

	fmt.Println(string(result))
	check(err)

	return string(result)
}

func readFile(filePath string) string {
	dat, err := os.ReadFile(filePath)
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

func createTask(name string, description string, timeEstimate int, priority PriorityLevel) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}

	return Task{name, description, timeEstimate, priority}, nil
}

func main() {

	testTask, err := createTask("Test", "Test task", 25, 1)

	check(err)

	writeTask(testTask)
}
