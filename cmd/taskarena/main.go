package main

import (
	"fmt"
	"os"

	"github.com/eduardo79silva/taskarena/internal/app"
	"github.com/eduardo79silva/taskarena/internal/cli"
	"github.com/eduardo79silva/taskarena/internal/config"
	"github.com/eduardo79silva/taskarena/internal/store"
)

func main() {
	cfg := loadConfig()

	s, err := store.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	a := app.New(cfg, s)
	root := cli.NewRootCmd(a)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func loadConfig() *config.Config {
	dir, err := config.GetDefaultConfigDir()
	if err != nil {
		return &config.DefaultConfig
	}
	conf, err := config.LoadConfig(dir + "/config.toml")
	if err != nil {
		return &config.DefaultConfig
	}
	return conf
}
