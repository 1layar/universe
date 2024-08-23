package dto

type (
	LoginDTO struct {
		Email    string `validate:"required,email,exist=email" json:"email"`
		Password string `validate:"required,min=6" json:"password"`
	}

	RegisterDTO struct {
		Email    string `validate:"required,email,exist=!email" json:"email"`
		Password string `validate:"required,min=6" json:"password"`
		Username string `validate:"required,min=5,max=20,exist=!username" json:"username"`
	}
)
