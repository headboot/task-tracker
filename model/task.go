package model

import (
	"math/rand"
	"time"
)

type Task struct {
	Id          uint64
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
		Id:          rand.Uint64(),
		Description: desc,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type UpdateOptions struct {
	Text  string
	Status Status
}

func (t *Task)Update(update UpdateOptions) Task { 
	var task Task
	if  update.Status == "" {
		task = Task{
			Id: t.Id,
			Description: update.Text,
			Status: Todo,
			CreatedAt: t.CreatedAt,
			UpdatedAt: time.Now(),
		}
	} else {
		task = Task{
			Id: t.Id,
			Description: t.Description,
			Status: update.Status,
			CreatedAt: t.CreatedAt,
			UpdatedAt: time.Now(),
		}
	}
	return task
}
