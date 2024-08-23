package model

import (
	"github.com/uptrace/bun"
)

type ProductCategoryRelation struct {
	bun.BaseModel `bun:"table:product_catalog.product_category_relations"`

	ProductID  int       `bun:",pk"`
	CategoryID int       `bun:",pk"`
	Product    *Product  `bun:"rel:belongs-to,join:product_id=id"`
	Category   *Category `bun:"rel:belongs-to,join:category_id=id"`
}
