package model

import (
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:ppob.products,alias:p"`

	ID          int    `bun:"id,pk,autoincrement"`
	Name        string `bun:"name,notnull"`
	SKU         string `bun:"sku,notnull,unique"`
	Description string `bun:"description,notnull"`
	PictureURL  string `bun:"picture_url,notnull"`
	Quantity    int    `bun:"quantity,notnull"`
	Price       string `bun:"price,notnull"`

	// many-to-many relationship with category
	Categories []*Category `bun:"m2m:ppob.product_category_relations,join:Product=Category"`
}
