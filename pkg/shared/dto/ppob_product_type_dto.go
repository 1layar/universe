package dto

type (
	PPobTypeDTO struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	GetAllPpobTypeResDTO struct {
		PaginateRespDto[[]PPobTypeDTO]
	}
)
