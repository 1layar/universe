package dto

type (
	AddProductDTO struct {
		Name        string `validate:"required" json:"name"`
		SKU         string `validate:"required,exist=!sku" json:"sku"`
		Description string `validate:"required" json:"description"`
		PictureURL  string `validate:"required" json:"picture_url"`
		Quantity    int    `validate:"required,min=1" json:"quantity"`
		Price       string `validate:"required" json:"price"`
		CategoryIDs []int  `validate:"required" json:"category_ids"` // many-to-many relationship with category_id
	}
	UpdateProductDTO struct {
		ID          int    `validate:"required" json:"id" swaggerignore:"true"`
		Name        string `validate:"required" json:"name"`
		SKU         string `validate:"required,exist=!sku:ID" json:"sku"`
		Description string `validate:"required" json:"description"`
		PictureURL  string `validate:"required" json:"picture_url"`
		Quantity    int    `validate:"required,min=1" json:"quantity"`
		Price       string `validate:"required" json:"price"`
		CategoryIDs []int  `validate:"required" json:"category_ids"` // many-to-many relationship with category_id
	}

	GetAllProductPaginateReqDTO struct {
		PaginateReqDto
		Name        string   `json:"name"`
		SKU         string   `json:"sku"`
		Description string   `json:"description"`
		Quantity    int      `json:"quantity"`
		Price       string   `json:"price"`
		Categories  []string `json:"categories"`
	}

	CategoryDTO struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	ProductDTO struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		SKU         string        `json:"sku"`
		Description string        `json:"description"`
		PictureURL  string        `json:"picture_url"`
		Quantity    int           `json:"quantity"`
		Price       string        `json:"price"`
		Categories  []CategoryDTO `json:"categories"`
	}

	GetAllProductPaginateResDTO struct {
		PaginateRespDto[[]ProductDTO]
	}
)
