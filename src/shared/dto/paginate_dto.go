package dto

type PaginateReqDto struct {
	Page  int `json:"page" query:"page"`
	Limit int `json:"limit" query:"limit"`
}

type PaginateRespDto[T any] struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
	Data  T     `json:"data"`
}
