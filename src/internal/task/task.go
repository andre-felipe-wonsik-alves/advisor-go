package task

import "time"

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	ReminderAt  time.Time `json:"reminder_at"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:created_at`
	UpdatedAt   time.Time `json:updated_at`
}
