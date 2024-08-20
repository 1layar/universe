package dto

type (
	CreateUserDto struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	SearchUserDto struct {
		Username *string `json:"username"`
	}

	UpdateUser struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	FindUser struct {
		Id int `json:"id"`
	}

	DeleteUser struct {
		Id int `json:"id"`
	}
)
