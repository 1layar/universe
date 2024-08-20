package model

import (
	"time"

	"github.com/uptrace/bun"
)

const (
	RoleUser = iota
	RoleAdmin
)

type User struct {
	bun.BaseModel `bun:"account.users" swaggerignore:"true"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Username  string    `bun:",unique,notnull" json:"username"`
	Email     string    `bun:",unique,notnull" json:"email"`
	Password  string    `bun:",notnull" json:"password"`
	Role      int       `bun:",notnull" json:"role"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
