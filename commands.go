package main

import (
	"flag"
	"fmt"
)

func runPush(args []string) {
	cmd := flag.NewFlagSet("push", flag.ExitOnError)

	var priority PriorityLevel = MediumPriority
	cmd.Var(&priority, "priority", "priority (low|medium|high|veryhigh)")
	cmd.Var(&priority, "p", "priority (shorthand)")

	var name string
	cmd.StringVar(&name, "name", "", "name")
	cmd.StringVar(&name, "n", "", "name (shorthand)")

	var description string
	cmd.StringVar(&description, "desc", "", "description")
	cmd.StringVar(&description, "d", "", "description (shorthand)")

	timeEstimate := cmd.Int("time", 25, "time estimate for the task")
	cmd.IntVar(timeEstimate, "t", 25, "time estimate (shorthand)")

	cmd.Parse(args)

	fmt.Println("subcommand 'push'")
	fmt.Println("  name:", name)

	task, err := createTask(name, description, *timeEstimate, priority)
	check(err)
	writeTask(task)
}

func runPull(args []string) {
	cmd := flag.NewFlagSet("pull", flag.ExitOnError)
	cmd.Parse(args)

	fmt.Println("subcommand 'pull'")
	tasks := loadTasks()
	task, err := selectNextTask(tasks)
	check(err)
	fmt.Println(task)
}
