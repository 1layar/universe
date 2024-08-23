package event

type ProductAddEvent struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}
