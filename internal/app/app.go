package app

import (
	"github.com/eduardo79silva/taskarena/internal/config"
	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/scheduler"
	"github.com/eduardo79silva/taskarena/internal/store"
)

func (a *App) DefaultPriority() priority.Level { return a.cfg.Defaults.Priority }

func (a *App) DefaultTimeEstimate() int { return a.cfg.Defaults.TimeEstimate }

type App struct {
	store     *store.Store
	scheduler *scheduler.Scheduler
	cfg       *config.Config
}

func New(cfg *config.Config, s *store.Store) *App {
	return &App{
		store:     s,
		scheduler: scheduler.New(cfg.Scheduler),
		cfg:       cfg,
	}
}
