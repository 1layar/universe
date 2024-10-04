package command

type (
	CartAddItemCommand struct {
		SessionId string
		ProductID int
		Quantity  int
	}

	CartAddItemCommandResult struct {
		Success bool
	}
)
