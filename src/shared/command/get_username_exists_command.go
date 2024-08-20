package command

type (
	GetUsernameExistsCommand struct {
		Username string
		Field    map[string]string
	}

	GetUsernameExistsResult struct {
		Exists bool
	}
)
