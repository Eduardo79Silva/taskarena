package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
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
	Name         string
	Description  string
	TimeEstimate int
	Priority     PriorityLevel
}

var TasksFilePath = "./tasks.json"

var Tasks []Task

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func appendTaskToJsonString(task Task, jsonText string) string {
	err := json.Unmarshal([]byte(jsonText), &Tasks)

	check(err)

	Tasks = append(Tasks, task)

	result, err := json.Marshal(Tasks)

	check(err)

	return string(result)
}

func readFile(filePath string) string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return string(dat)
}

func writeTask(task Task, fileHandle *os.File) {
	jsonText := readFile(TasksFilePath)

	if jsonText == "" {
		jsonText = "{}"
	}

	parsedJson := appendTaskToJsonString(task, jsonText)

	writer := bufio.NewWriter(fileHandle)

	_, err := writer.WriteString(parsedJson)

	check(err)
}

func checkFileExists(filePath string) *os.File {
	fileHandle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return fileHandle
}

func createTask(name string, description string, timeEstimate int, priority PriorityLevel) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}

	return Task{name, description, timeEstimate, priority}, nil
}

func main() {

	testTask := Task{"Test", "Test task", 25, 1}

	fileHandle := checkFileExists(TasksFilePath)

	writeTask(testTask, fileHandle)

}
