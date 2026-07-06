package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type command struct {
	name string
	run  func(args []string)
}

var commands = []command{
	{"push", runPush},
	{"pull", runPull},
	{"done", runDone},
	{"status", runStatus},
	{"list", runList},
}

var AppConfig = loadAppConfig()

func loadAppConfig() Config {
	configDir, err := GetDefaultConfigDir()
	if err != nil {
		return DefaultConfig
	}

	conf, err := LoadConfig(filepath.Join(configDir, "config.toml"))
	if err != nil {
		return DefaultConfig
	}

	return *conf
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'push' or 'pull' subcommands")
		os.Exit(1)
	}

	for _, c := range commands {
		if c.name == os.Args[1] {
			c.run(os.Args[2:])
			return
		}
	}

	fmt.Println("expected 'push' or 'pull' subcommands")

	os.Exit(1)
}
