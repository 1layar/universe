package event

type UserRegisterEvent struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	ConfirmLink string `json:"confirm_link"`
	ConfirmCode string `json:"confirm_code"`
	Password    string `json:"password"`
}
