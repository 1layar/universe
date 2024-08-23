package command

type (
	SearchUserCommand struct {
		Username *string `json:"string"`
	}

	SearchUserResult struct {
		Users []AddUserResult `json:"users"`
	}
)
