package command

type (
	GetUserCommand struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	GetUserResult struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
