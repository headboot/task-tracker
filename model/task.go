package model

import (
	"math/rand"
	"time"
)

type Task struct {
	id          uint
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

// statuses of Task
type Status string

const (
	Todo     Status = "todo"
	Progress Status = "progress"
	Done     Status = "done"
)

func New(desc string) Task {
	return Task{
		id:          uint(rand.Uint64()),
		description: desc,
		status:      Todo,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}
