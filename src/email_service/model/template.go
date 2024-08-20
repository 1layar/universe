package model

import (
	"github.com/uptrace/bun"
)

// EmailTemplate represents the schema for the email.templates table
type EmailTemplate struct {
	bun.BaseModel `bun:"table:email.templates,alias:t"`

	ID           int            `bun:"id,pk,autoincrement"`
	Name         string         `bun:"name,notnull"`
	Subject      string         `bun:"subject,notnull"`
	TextContent  string         `bun:"text_content,nullzero"`
	HtmlContent  string         `bun:"html_content,nullzero"`
	Placeholders map[string]any `bun:"placeholders,nullzero"`

	Events []*EmailEvent `bun:"rel:has-many,join:id=template_id"`
}
