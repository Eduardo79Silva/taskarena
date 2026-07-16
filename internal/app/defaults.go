package app

import (
	"github.com/eduardo79silva/taskarena/internal/priority"
)

func (a *App) DefaultPriority() priority.Level { return a.cfg.Defaults.Priority }

func (a *App) DefaultTimeEstimate() int { return a.cfg.Defaults.TimeEstimate }
