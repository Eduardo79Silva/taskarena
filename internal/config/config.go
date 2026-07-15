package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/mitchellh/go-homedir"
)

var configDirName = "taskarena"

func GetDefaultConfigDir() (string, error) {
	var configDirLocation string

	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "linux":
		// Use the XDG_CONFIG_HOME variable if it is set, otherwise
		// $HOME/.config/taskarena
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			configDirLocation = xdgConfigHome
		} else {
			configDirLocation = filepath.Join(homeDir, ".config", configDirName)
		}

	default:
		// On other platforms we just use $HOME/.taskarena
		hiddenConfigDirName := "." + configDirName
		configDirLocation = filepath.Join(homeDir, hiddenConfigDirName)
	}

	return configDirLocation, nil
}

type Config struct {
	Scheduler SchedulerConfig
	Defaults  TaskConfig
	Status    StatusConfig
}

type SchedulerConfig struct {
	PriorityWeight     float64
	TimeWeight         float64
	AgingWeight        float64
	SelectionSharpness float64
	AgingHorizonHours  float64
}

type TaskConfig struct {
	Priority     priority.Level
	TimeEstimate int
}

type StatusConfig struct {
	Format string
}

var DefaultConfig = Config{
	Scheduler: SchedulerConfig{
		PriorityWeight:     0.5,
		TimeWeight:         0.2,
		AgingWeight:        0.3,
		SelectionSharpness: 2.15,
		AgingHorizonHours:  7 * 24.0,
	},
	Defaults: TaskConfig{
		Priority:     priority.Medium,
		TimeEstimate: 25,
	},
	Status: StatusConfig{
		Format: "plain",
	},
}

func LoadConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("config file does not exist")
	} else if err != nil {
		return nil, err
	}

	conf := DefaultConfig
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
