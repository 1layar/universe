package model

import (
	"time"

	"github.com/uptrace/bun"
)

type ProductKind string

const (
	KindPrepaid  ProductKind = "prepaid"
	KindPostpaid ProductKind = "postpaid"
)

type ProductStatus string

const (
	StatusActive   ProductStatus = "active"
	StatusInactive ProductStatus = "inactive"
)

type Product struct {
	bun.BaseModel `bun:"table:ppob.products,alias:p"`

	ID           int           `bun:"id,pk,autoincrement"`
	Kind         ProductKind   `bun:"kind,notnull"`
	Code         string        `bun:"product_code,notnull,unique"`
	Description  string        `bun:"product_description,notnull"`
	Nominal      string        `bun:"product_nominal,notnull"`
	Details      string        `bun:"product_details"`
	Price        float64       `bun:"product_price,notnull"`
	TypeId       int           `bun:"product_type_id,notnull"`
	ActivePeriod *int          `bun:"active_period"`
	Status       ProductStatus `bun:"status,notnull"`
	IconURL      string        `bun:"icon_url"`
	CategoryId   int           `bun:"product_category_id,notnull"`
	BillingCycle *int          `bun:"billing_cycle"`
	DueDate      *time.Time    `bun:"due_date"`
	GracePeriod  *int          `bun:"grace_period"`
	Comission    float64       `bun:"comission"`
	Fee          float64       `bun:"fee"`
	// many-to-many relationship with category
	Category *ProductCategory `bun:"rel:belongs-to,join:product_category_id=id"`

	// many-to-many relationship with product type
	ProductType *ProductType `bun:"rel:belongs-to,join:product_type_id=id"`
}
