package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:product_catalog.categories,alias:c"`

	ID          int     `bun:"id,pk,autoincrement"`
	Name        string  `bun:"name"`
	Description *string `bun:"description,nullzero"`
	PictureURL  *string `bun:"picture_url,nullzero"`
	ParentID    *int    `bun:"parent_id"`

	Parent   *Category   `bun:"rel:belongs-to,join:parent_id=id"`
	Children []*Category `bun:"rel:has-many,join:id=parent_id"`
	Products []*Product  `bun:"m2m:product_catalog.product_category_relations,join:Category=Product"`
}
