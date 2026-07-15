package scheduler

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/eduardo79silva/taskarena/internal/config"
	"github.com/eduardo79silva/taskarena/internal/priority"
	"github.com/eduardo79silva/taskarena/internal/task"
)

type Scheduler struct {
	cfg config.SchedulerConfig
}

func New(cfg config.SchedulerConfig) *Scheduler {
	return &Scheduler{
		cfg: cfg,
	}
}

var errEmptyTaskList = errors.New("no tasks to pull from")

func filterTasksByTag(tasks []task.Task, tag string) []task.Task {
	var filteredTasks []task.Task

	for _, task := range tasks {
		if task.Tag == tag {
			filteredTasks = append(filteredTasks, task)

		}
	}

	return filteredTasks
}

func filterTasksByTime(tasks []task.Task, timeLimit int) []task.Task {
	var filteredTasks []task.Task

	for _, task := range tasks {
		if task.TimeEstimate <= timeLimit {
			filteredTasks = append(filteredTasks, task)

		}
	}

	return filteredTasks
}

func (s *Scheduler) wsmScore(t task.Task, minTime, maxTime int) float64 {
	normPriority := float64(t.Priority) / float64(priority.VeryHigh)

	normTime := 1.0
	if maxTime > minTime {
		normTime = 1.0 - float64(t.TimeEstimate-minTime)/float64(maxTime-minTime)
	}

	age := time.Since(t.CreatedAt).Hours()

	normAge := min(age/s.cfg.AgingHorizonHours, 1.0)

	return s.cfg.PriorityWeight*normPriority + s.cfg.TimeWeight*normTime + s.cfg.AgingWeight*normAge
}

func (s *Scheduler) SelectNextTask(tasks []task.Task) (task.Task, error) {
	if len(tasks) == 0 {
		return task.Task{}, errEmptyTaskList
	}

	minTime, maxTime := tasks[0].TimeEstimate, tasks[0].TimeEstimate
	for _, t := range tasks[1:] {
		if t.TimeEstimate < minTime {
			minTime = t.TimeEstimate
		}
		if t.TimeEstimate > maxTime {
			maxTime = t.TimeEstimate
		}
	}

	scores := make([]float64, len(tasks))
	total := 0.0
	for i, t := range tasks {
		scores[i] = math.Pow(s.wsmScore(t, minTime, maxTime), s.cfg.SelectionSharpness)
		total += scores[i]
	}

	if total == 0 {
		return tasks[rand.Intn(len(tasks))], nil
	}

	r := rand.Float64() * total
	cumulative := 0.0
	for i, s := range scores {
		cumulative += s
		if r < cumulative {
			return tasks[i], nil
		}
	}

	return tasks[len(tasks)-1], nil
}
