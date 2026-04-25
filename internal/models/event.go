package models

import "time"

type Event struct {
	ID         int64          `json:"id"`
	UserID     string         `json:"user_id"`
	EventType  string         `json:"event_type"`
	Timestamp  time.Time      `json:"timestamp"`
	Properties map[string]any `json:"properties"`
}
