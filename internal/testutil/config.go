package testutil

import "github.com/eduardo79silva/taskarena/internal/config"

func MakeSchedulerConfig() config.SchedulerConfig {
	return config.SchedulerConfig{
		AgingHorizonHours:  168,
		PriorityWeight:     1.0,
		TimeWeight:         1.0,
		AgingWeight:        1.0,
		SelectionSharpness: 2.0,
	}
}
