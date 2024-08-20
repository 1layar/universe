package dto

import "time"

type Compose struct {
	ID        string
	Event     string
	Agent     string
	Email     string
	Template  string
	EventTime time.Time
	Payload   map[string]any
}
