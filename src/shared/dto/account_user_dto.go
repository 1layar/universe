package dto

type (
	AddUser struct {
		Username string `validate:"required,min=5,max=20,exist=!username" json:"username"`
		Email    string `validate:"required,email,exist=!email" json:"email"`
		Password string `validate:"required,min=6" json:"password"`
		Role     int    `validate:"required,oneof=0 1" json:"role"`
	}

	UpdateUser struct {
		Id       int    `validate:"required" json:"id"`
		Username string `validate:"required,min=5,max=20,exist=!username:Id" json:"username"`
		Email    string `validate:"required,email,exist=!email:Id" json:"email"`
		Password string `validate:"required,min=6" json:"password"`
		Role     int    `validate:"required,oneof=0 1" json:"role"`
	}

	SearchUser struct {
		Username *string `json:"username" query:"username"`
	}

	FindUser struct {
		ID int `validate:"required"`
	}
)
