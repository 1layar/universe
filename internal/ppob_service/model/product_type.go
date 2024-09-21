package model

import "github.com/uptrace/bun"

type ProductType struct {
	bun.BaseModel `bun:"table:ppob.product_types,alias:t"`

	ID       int        `bun:"id,pk,autoincrement"`
	Name     string     `bun:"type_name"`
	Products []*Product `bun:"rel:has-many,join:id=product_type_id"`
}
