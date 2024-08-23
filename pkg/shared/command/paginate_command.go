package command

type (
	PaginateCommand struct {
		Page  int `json:"page" query:"page"`
		Limit int `json:"limit" query:"limit"`
	}

	PaginateResponse[T any] struct {
		Total int `json:"total"`
		Data  T   `json:"data"`
	}
)
