package models

import "time"

type Todo struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Priority    Priority   `json:"priority"`
	Status      Status     `json:"status"`
	Completed   bool       `json:"completed"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
	CompletedAt *time.Time `json:"completedAt"`
}
type Status string
type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
	PriorityNone   Priority = "none"
)

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)
