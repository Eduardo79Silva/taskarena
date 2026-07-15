package store

import (
	"encoding/json"
	"github.com/eduardo79silva/taskarena/internal/task"
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

func ReadTasksFile(filePath string) ([]task.Task, error) {
	dat, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(dat) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task
	err = json.Unmarshal(dat, &tasks)
	return tasks, err
}

func ReadTaskFile(filePath string) (task.Task, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return task.Task{}, err
	}
	var task task.Task
	err = json.Unmarshal(dat, &task)
	return task, err
}

func WriteTasksFile(filePath string, tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func WriteTaskFile(filePath string, task task.Task) error {
	data, err := json.MarshalIndent(task, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func PushTask(filePath string, task task.Task) {
	tasks, err := ReadTasksFile(filePath)
	check(err)
	tasks = append(tasks, task)
	check(WriteTasksFile(filePath, tasks))
}

func WriteAllTasks(tasks []task.Task) {
	check(WriteTasksFile(TasksFilePath, tasks))
}

func WriteCurrentTask(task task.Task) {
	check(WriteTaskFile(CurrentTaskFilePath, task))
}

func DeleteTask(tasks []task.Task, taskID string) []task.Task {
	for i, task := range tasks {
		if task.ID == taskID {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func DeleteTaskFromFile(filePath string, taskID string) {
	tasks, err := ReadTasksFile(filePath)
	check(err)
	tasks = DeleteTask(tasks, taskID)
	check(WriteTasksFile(filePath, tasks))
}

func ClearCurrentTask() error {
	err := os.Remove(CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func CompleteCurrentTask() {
	current, err := ReadTaskFile(CurrentTaskFilePath)
	check(err)

	now := time.Now()
	current.CompletedAt = &now

	calculateTimeSpent(&current)
	updateTaskTime(&current)

	PushTask(CompletedTasksFilePath, current)

	check(ClearCurrentTask())
}

func LoadTasks() []task.Task {
	tasks, err := ReadTasksFile(TasksFilePath)
	check(err)
	return tasks
}
