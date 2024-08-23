package command

type (
	LoginCommand struct {
		Email    string
		Password string
		Ip       string
		Ua       string
		DevId    string
		Retry    string
	}

	LoginResult struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiredAt    string `json:"expired_at"`
		Permission   []string
	}
)
