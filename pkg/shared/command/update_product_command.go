package command

type (
	UpdateProductCommand struct {
		ID          int
		UserID      int
		Name        string
		SKU         string
		Description string
		PictureURL  string
		Quantity    int
		Price       string

		// many
		Categories []int
	}

	UpdateProductResult struct {
		ID int
	}
)
