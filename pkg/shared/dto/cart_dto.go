package dto

type (
	AddCart struct {
		ProductID int `validate:"required" json:"product_id"`
		Quantity  int `validate:"required" json:"quantity"`
	}

	UpdateCart struct {
		ProductID int `validate:"required" json:"product_id"`
		Quantity  int `validate:"required" json:"quantity"`
	}

	CartItem struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	Cart struct {
		Items []CartItem `json:"items"`
	}

	CartResp struct {
		SessionId string `json:"session_id"`
		Cart      Cart   `json:"cart"`
	}
)
