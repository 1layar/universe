package command

type (
	AddProductCommand struct {
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

	AddProductResult struct {
		ID int
	}
)
