package main

import (
	"fmt"
	"os"
)

type command struct {
	name string
	run  func(args []string)
}

var commands = []command{
	{"push", runPush},
	{"pull", runPull},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'push' or 'pull' subcommands")
		os.Exit(1)
	}

	counts := map[string]int{}
	for range 10000 {
		t, _ := selectNextTask(Tasks)
		counts[t.Name+t.Description]++ // disambiguate the two "Yay" tasks
	}
	fmt.Println(counts)

	for _, c := range commands {
		if c.name == os.Args[1] {
			c.run(os.Args[2:])
			return
		}
	}

	fmt.Println("expected 'push' or 'pull' subcommands")

	os.Exit(1)
}
