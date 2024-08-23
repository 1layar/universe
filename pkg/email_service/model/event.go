package model

import (
	"time"

	"github.com/uptrace/bun"
)

type EmailEvent struct {
	bun.BaseModel `bun:"table:email.events,alias:e"`

	ID         int            `bun:"id,pk,autoincrement"`
	MessageID  int            `bun:"message_id,notnull"`
	TemplateID int            `bun:"template_id,notnull"`
	EventType  string         `bun:"event_type,notnull"`
	Payload    map[string]any `bun:"payload,nullzero"`
	EventTime  time.Time      `bun:"event_time,notnull,default:current_timestamp"`

	Message  *EmailMessage  `bun:"rel:belongs-to,join:message_id=id"`
	Template *EmailTemplate `bun:"rel:belongs-to,join:template_id=id"`
}
