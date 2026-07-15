package actions

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
)

func runPush(args []string) {
	cmd := flag.NewFlagSet("push", flag.ExitOnError)

	var priority priority.Level = AppConfig.Defaults.Priority
	cmd.Var(&priority, "priority", "priority (low|medium|high|veryhigh)")
	cmd.Var(&priority, "p", "priority (shorthand)")

	var name string
	cmd.StringVar(&name, "name", "", "name")
	cmd.StringVar(&name, "n", "", "name (shorthand)")

	var description string
	cmd.StringVar(&description, "desc", "", "description")
	cmd.StringVar(&description, "d", "", "description (shorthand)")

	timeEstimate := cmd.Int("time", AppConfig.Defaults.TimeEstimate, "time estimate for the task")
	cmd.IntVar(timeEstimate, "t", AppConfig.Defaults.TimeEstimate, "time estimate (shorthand)")

	var tag string
	cmd.StringVar(&tag, "tag", "", "tag")

	cmd.Parse(args)

	fmt.Println("subcommand 'push'")
	fmt.Println("  name:", name)

	task, err := createTask(name, description, *timeEstimate, priority, tag)
	check(err)
	pushTask(TasksFilePath, task)
}

func runPull(args []string) {
	cmd := flag.NewFlagSet("pull", flag.ExitOnError)

	tagFilter := cmd.String("tag", "", "filter by tag")
	timeFilter := cmd.Int("time", -1, "filter by time")

	cmd.Parse(args)

	fmt.Println("subcommand 'pull'")
	tasks := loadTasks()

	if *tagFilter != "" {
		tasks = filterTasksByTag(tasks, *tagFilter)
	}

	if *timeFilter != -1 {
		tasks = filterTasksByTime(tasks, *timeFilter)
	}

	newTask, err := selectNextTask(tasks)
	updateTaskTime(&newTask)

	check(err)

	currentTask, _ := readTaskFile(CurrentTaskFilePath)
	updateTaskTime(&currentTask)
	calculateTimeSpent(&currentTask)

	if (currentTask != task.Task{}) {
		pushTask(TasksFilePath, currentTask)
		tasks = loadTasks()
	} else if currentTask.ID == newTask.ID {
		fmt.Println(newTask)
	}

	writeCurrentTask(newTask)
	tasks = deleteTask(tasks, newTask.ID)
	writeAllTasks(tasks)

	fmt.Println(newTask)
}

func runDone(args []string) {
	cmd := flag.NewFlagSet("done", flag.ExitOnError)
	cmd.Parse(args)

	fmt.Println("subcommand 'done'")
	completeCurrentTask()
}

func runStatus(args []string) {
	cmd := flag.NewFlagSet("status", flag.ExitOnError)

	format := cmd.String("format", "plain", "output format (plain|waybar|notify)")
	cmd.Parse(args)

	view, err := getCurrentTaskView()
	check(err)

	switch *format {
	case "waybar":
		out, err := formatWaybar(view)
		check(err)
		fmt.Println(out)
	case "notify":
		fmt.Println(formatNotification(view))
	case "plain":
		fmt.Println(formatPlain(view))
	default:
		fmt.Println("unknown format:", *format)
		os.Exit(1)
	}
}

func runList(args []string) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	cmd.Parse(args)

	tasks := loadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "TAG\tPRIORITY\tTIME\tNAME")
	for _, t := range tasks {
		fmt.Fprintf(w, "%s\t%s\t%dm\t%s\n", t.Tag, t.Priority.String(), t.TimeEstimate, t.Name)
	}
	w.Flush()
}

func runEdit(args []string) {
	cmd := flag.NewFlagSet("edit", flag.ExitOnError)
	cmd.Parse(args)

	if err := EditFile(TasksFilePath); err != nil {
		panic(err)
	}

}
