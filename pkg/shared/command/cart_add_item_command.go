package command

type (
	CartAddItemCommand struct {
		SessionID string
		ProductID int
		Quantity  int
		Source    string
	}

	CartAddItemCommandResult struct {
		SessionID string
		Success   bool
	}
)
