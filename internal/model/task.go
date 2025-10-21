package model

import (
	"time"
)

const FILENAME = "tasks.json"

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      string(TASK_STATUS_TODO),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
