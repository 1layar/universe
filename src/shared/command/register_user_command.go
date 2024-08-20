package command

type (
	RegisterCommand struct {
		Username string
		Email    string
		Password string
		Ip       string
		Ua       string
		DevId    string
		Retry    string
	}

	RegisterResult struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiredAt    string `json:"expired_at"`
		Permission   []string
	}
)
