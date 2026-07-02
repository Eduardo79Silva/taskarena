package main

import (
	"encoding/json"
	"os"
)

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
	result, err := json.MarshalIndent(Tasks, "", "\t")
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

func pullTask() Task {
	jsonText := readFile(TasksFilePath)
	if jsonText == "" {
		return Task{}
	}
	err := json.Unmarshal([]byte(jsonText), &Tasks)
	check(err)
	return Tasks[0]
}
