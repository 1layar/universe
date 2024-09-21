package model

import "github.com/uptrace/bun"

type ProductCategory struct {
	bun.BaseModel `bun:"table:ppob.product_categories,alias:c"`

	ID       int        `bun:"id,pk,autoincrement"`
	Name     string     `bun:"category_name"`
	Products []*Product `bun:"rel:has-many,join:id=product_category_id"`
}
