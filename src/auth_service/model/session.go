package model

import (
	"time"

	"github.com/uptrace/bun"
)

type SessionKind string

const (
	LoginKind    SessionKind = "LoginKind"
	RefereshKind SessionKind = "RefereshKind"
	LogoutKind   SessionKind = "LogoutKind"
	IddleKind    SessionKind = "IddleKind"
)

type Session struct {
	bun.BaseModel `bun:"auth.sessions" swaggerignore:"true"`

	ID        int         `bun:",pk,autoincrement" json:"id"`
	UserID    int         `bun:",notnull" json:"user_id"`
	IP        string      `bun:",notnull" json:"ip"`
	UserAgent string      `bun:",notnull" json:"user_agent"`
	Kind      SessionKind `bun:",notnull" json:"kind"`
	Retry     int         `bun:",notnull" json:"retry"`
	CreatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
}
