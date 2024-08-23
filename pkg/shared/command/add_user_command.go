package command

type (
	AddUserCommand struct {
		Username string
		Email    string
		Password string
		Role     int
	}

	AddUserResult struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Role     int    `json:"role"`
		Password string `json:"password"`
	}
)
