package model

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	SessionId string     `json:"session_id"`
	Items     []CartItem `json:"items"`
}
