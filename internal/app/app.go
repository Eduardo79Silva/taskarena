package app

import (
	"github.com/eduardo79silva/taskarena/internal/config"
	"github.com/eduardo79silva/taskarena/internal/scheduler"
	"github.com/eduardo79silva/taskarena/internal/store"
)

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
