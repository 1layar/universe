package command

type (
	GetAllProductCommand struct {
		PaginateCommand
		Name        string   `json:"name"`
		SKU         string   `json:"sku"`
		Description string   `json:"description"`
		Quantity    int      `json:"quantity"`
		Price       string   `json:"price"`
		Categories  []string `json:"categories"`
	}

	CategoryResult struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	ProductResult struct {
		ID          int              `json:"id"`
		Name        string           `json:"name"`
		SKU         string           `json:"sku"`
		Description string           `json:"description"`
		PictureURL  string           `json:"picture_url"`
		Quantity    int              `json:"quantity"`
		Price       string           `json:"price"`
		Categories  []CategoryResult `json:"categories"`
	}

	GetAllProductResult struct {
		PaginateResponse[[]ProductResult]
	}
)
