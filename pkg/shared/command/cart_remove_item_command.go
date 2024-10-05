package command

type (
	CartRemoveItemCommand struct {
		SessionID string
		ProductID int
		Quantity  int
		Source    string
	}

	CartRemoveItemCommandResult struct {
		SessionID string
		Success   bool
	}
)
