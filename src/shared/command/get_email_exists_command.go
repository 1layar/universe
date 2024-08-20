package command

type (
	GetEmailExistsCommand struct {
		Email string
		Field map[string]string
	}

	GetEmailExistsResult struct {
		Exists bool
	}
)
