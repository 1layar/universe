package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:email.accounts,alias:a"`

	ID           int       `bun:"id,pk,autoincrement"`
	Code         string    `bun:"code,unique,notnull"`
	SMTPHost     string    `bun:"smtp_host,notnull"`
	SMTPPort     int       `bun:"smtp_port,notnull"`
	SMTPUsername string    `bun:"smtp_username,notnull"`
	SMTPPassword string    `bun:"smtp_password,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Outbox []*EmailMessage `bun:"rel:has-many,join:id=account_id"`
}
