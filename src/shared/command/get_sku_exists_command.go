package command

type (
	GetSKUExistsCommand struct {
		SKU   string
		Field map[string]string
	}

	GetSKUExistsResult struct {
		Exists bool
	}
)
