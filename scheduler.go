package main

import (
	"errors"
	"math"
	"math/rand"
)

var errEmptyTaskList = errors.New("no tasks to pull from")

const (
	priorityWeight = 0.7
	timeWeight     = 0.3
)

const selectionSharpness = 2.15

func wsmScore(t Task, minTime, maxTime int) float64 {
	normPriority := float64(t.Priority) / float64(VeryHighPriority)

	normTime := 1.0
	if maxTime > minTime {
		normTime = 1.0 - float64(t.TimeEstimate-minTime)/float64(maxTime-minTime)
	}

	return priorityWeight*normPriority + timeWeight*normTime
}

func selectNextTask(tasks []Task) (Task, error) {
	if len(tasks) == 0 {
		return Task{}, errEmptyTaskList
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
		scores[i] = math.Pow(wsmScore(t, minTime, maxTime), selectionSharpness)
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
