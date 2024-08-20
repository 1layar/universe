package model

import (
	"time"

	"github.com/uptrace/bun"
)

type EmailMessage struct {
	bun.BaseModel `bun:"table:email.messages,alias:m" swaggerignore:"true"`

	ID        int       `bun:"id,pk,autoincrement"`
	AccountID int       `bun:"account_id,notnull"`
	ToEmail   string    `bun:"to_email,notnull"`
	Subject   string    `bun:"subject,notnull"`
	TextBody  string    `bun:"text_body,nullzero"`
	HtmlBody  string    `bun:"html_body,nullzero"`
	Status    string    `bun:"status,notnull,default:'pending'"`
	SendAt    time.Time `bun:"send_at,nullzero"`
	OpenAt    time.Time `bun:"open_at,nullzero"`
	ReadAt    time.Time `bun:"read_at,nullzero"`
	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp"`

	Events  []*EmailEvent `bun:"rel:has-many,join:id=message_id"`
	Account *Account      `bun:"rel:belongs-to,join:account_id=id"`
}
