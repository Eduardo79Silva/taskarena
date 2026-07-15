package priority

import (
	"fmt"
	"strings"
)

type Level int

const (
	Low Level = iota
	Medium
	High
	VeryHigh
)

func (p *Level) String() string {
	return [...]string{"low", "medium", "high", "veryhigh"}[*p]
}

func (p *Level) Set(s string) error {
	switch strings.ToLower(s) {
	case "low":
		*p = Low
	case "medium":
		*p = Medium
	case "high":
		*p = High
	case "veryhigh":
		*p = VeryHigh
	default:
		return fmt.Errorf("invalid priority: %s", s)
	}
	return nil
}
