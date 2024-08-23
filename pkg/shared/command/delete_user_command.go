package command

type (
	DeleteUserCommand struct {
		Id int `json:"id"`
	}

	DeleteUserResult struct {
		Id        int  `json:"id"`
		HasDelete bool `json:"has_delete"`
	}
)
