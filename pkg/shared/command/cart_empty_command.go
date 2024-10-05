package command

type (
	CartEmptyCommand struct {
		SessionId string
		Source    string
	}

	CartEmptyCommandResult struct {
		Success bool
	}
)
