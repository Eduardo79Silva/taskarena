package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/eduardo79silva/taskarena/internal/task"
)

type Store struct {
	TasksFilePath          string
	CurrentTaskFilePath    string
	CompletedTasksFilePath string
}

func New() (*Store, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	dir := filepath.Join(home, ".config", "taskarena")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &Store{
		TasksFilePath:          filepath.Join(dir, "tasks.json"),
		CurrentTaskFilePath:    filepath.Join(dir, "current.json"),
		CompletedTasksFilePath: filepath.Join(dir, "completed.json"),
	}, nil
}

func (s *Store) LoadTasks() ([]task.Task, error) {
	return readTasksFile(s.TasksFilePath)
}

func (s *Store) ReadCurrentTask() (task.Task, error) {
	return readTaskFile(s.CurrentTaskFilePath)
}

func (s *Store) PushTask(t task.Task) error {
	tasks, err := readTasksFile(s.TasksFilePath)
	if err != nil {
		return err
	}
	tasks = append(tasks, t)
	return writeTasksFile(s.TasksFilePath, tasks)
}

func (s *Store) WriteAllTasks(tasks []task.Task) error {
	return writeTasksFile(s.TasksFilePath, tasks)
}

func (s *Store) WriteCurrentTask(t task.Task) error {
	return writeTaskFile(s.CurrentTaskFilePath, t)
}

func (s *Store) ClearCurrentTask() error {
	err := os.Remove(s.CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func (s *Store) CompleteCurrentTask() error {
	current, err := readTaskFile(s.CurrentTaskFilePath)
	if err != nil {
		return err
	}

	current.UpdateTime()
	current.CalculateTimeSpent()

	now := time.Now()
	current.CompletedAt = &now

	tasks, err := readTasksFile(s.CompletedTasksFilePath)
	if err != nil {
		return err
	}

	tasks = append(tasks, current)
	if err := writeTasksFile(s.CompletedTasksFilePath, tasks); err != nil {
		return err
	}

	return s.ClearCurrentTask()
}

func (s *Store) GetCurrentTaskView() (*task.CurrentTaskView, error) {
	t, err := readTaskFile(s.CurrentTaskFilePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task.CurrentTaskView{
		Name:         t.Name,
		Description:  t.Description,
		Priority:     t.Priority,
		TimeEstimate: t.TimeEstimate,
	}, nil
}

func readTasksFile(path string) ([]task.Task, error) {
	dat, err := os.ReadFile(path)
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

func readTaskFile(path string) (task.Task, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return task.Task{}, err
	}
	var t task.Task
	err = json.Unmarshal(dat, &t)
	return t, err
}

func writeTasksFile(path string, tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func writeTaskFile(path string, t task.Task) error {
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
