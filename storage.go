package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

var TasksFilePath, CurrentTaskFilePath, CompletedTasksFilePath = defaultFilesPath()

func defaultFilesPath() (string, string, string) {
	home, err := os.UserHomeDir()
	check(err)
	dir := filepath.Join(home, ".config", "taskarena")
	err = os.MkdirAll(dir, 0755)
	check(err)
	return filepath.Join(dir, "tasks.json"), filepath.Join(dir, "current.json"), filepath.Join(dir, "completed.json")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readTasksFile(filePath string) ([]Task, error) {
	dat, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(dat) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(dat, &tasks)
	return tasks, err
}

func readTaskFile(filePath string) (Task, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return Task{}, err
	}
	var task Task
	err = json.Unmarshal(dat, &task)
	return task, err
}

func writeTasksFile(filePath string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func writeTaskFile(filePath string, task Task) error {
	data, err := json.MarshalIndent(task, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func pushTask(filePath string, task Task) {
	tasks, err := readTasksFile(filePath)
	check(err)
	tasks = append(tasks, task)
	check(writeTasksFile(filePath, tasks))
}

func writeAllTasks(tasks []Task) {
	check(writeTasksFile(TasksFilePath, tasks))
}

func writeCurrentTask(task Task) {
	check(writeTaskFile(CurrentTaskFilePath, task))
}

func deleteTask(tasks []Task, taskID string) []Task {
	for i, task := range tasks {
		if task.ID == taskID {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func deleteTaskFromFile(filePath string, taskID string) {
	tasks, err := readTasksFile(filePath)
	check(err)
	tasks = deleteTask(tasks, taskID)
	check(writeTasksFile(filePath, tasks))
}

func clearCurrentTask() error {
	err := os.Remove(CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func completeCurrentTask() {
	current, err := readTaskFile(CurrentTaskFilePath)
	check(err)

	now := time.Now()
	current.CompletedAt = &now

	pushTask(CompletedTasksFilePath, current)

	check(clearCurrentTask())
}

func loadTasks() []Task {
	tasks, err := readTasksFile(TasksFilePath)
	check(err)
	return tasks
}
