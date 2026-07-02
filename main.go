package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type PriorityLevel int

const (
	LowPriority PriorityLevel = iota
	MediumPriority
	HighPriority
	VeryHighPriority
)

func (p *PriorityLevel) String() string {
	return [...]string{"low", "medium", "high", "veryhigh"}[*p]
}

func (p *PriorityLevel) Set(s string) error {
	switch strings.ToLower(s) {
	case "low":
		*p = LowPriority
	case "medium":
		*p = MediumPriority
	case "high":
		*p = HighPriority
	case "veryhigh":
		*p = VeryHighPriority
	default:
		return fmt.Errorf("invalid priority: %s", s)
	}
	return nil
}

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

func createTask(name string, description string, timeEstimate int, priority PriorityLevel) (Task, error) {
	if name == "" {
		return Task{}, errors.New("empty name")
	}

	return Task{name, description, timeEstimate, priority}, nil
}

func main() {

	pushCmd := flag.NewFlagSet("push", flag.ExitOnError)

	var pushTaskPriority PriorityLevel = MediumPriority
	pushCmd.Var(&pushTaskPriority, "priority", "priority (low|medium|high|veryhigh)")
	pushCmd.Var(&pushTaskPriority, "p", "priority (shorthand)")

	var pushTaskName string
	pushCmd.StringVar(&pushTaskName, "name", "", "name")
	pushCmd.StringVar(&pushTaskName, "n", "", "name (shorthand)")

	var pushTaskDescription string
	pushCmd.StringVar(&pushTaskDescription, "desc", "", "description")
	pushCmd.StringVar(&pushTaskDescription, "d", "", "description (shorthand)")

	pushTaskTimeEstimate := pushCmd.Int("time", 25, "time estimate for the task")
	pushCmd.IntVar(pushTaskTimeEstimate, "t", 25, "time estimate (shorthand)")

	pullCmd := flag.NewFlagSet("pull", flag.ExitOnError)

	switch os.Args[1] {

	case "push":
		pushCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'push'")
		fmt.Println("  name:", pushTaskName)
		testTask, err := createTask(pushTaskName, pushTaskDescription, *pushTaskTimeEstimate, PriorityLevel(pushTaskPriority))
		check(err)

		writeTask(testTask)
	case "pull":
		pullCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'pull'")

		task := pullTask()

		fmt.Println(task)
	default:
		fmt.Println("expected 'push' or 'pull' subcommands")
		os.Exit(1)
	}

}
