package models

import "time"

type Todo struct {
	Title       string
	Description string
	Category    string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
