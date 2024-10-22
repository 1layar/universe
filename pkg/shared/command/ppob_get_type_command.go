package command

type (
	PpobGetTypeCommand struct {
		PaginateCommand
	}

	PpobTypeResult struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	PpobGetTypeResult struct {
		PaginateResponse[[]PpobTypeResult]
	}
)
