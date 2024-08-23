package command

type JwtToUserCommand struct {
	Jwt string
}

type JwtToUserResponse struct {
	User GetUserResult
}
