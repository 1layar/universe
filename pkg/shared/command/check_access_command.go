package command

type (
	CheckAccessCommand struct {
		UserId     int
		Permission string
		Subject    string
		SubjectId  int
	}

	CheckAccessResponse struct {
		Allow bool
	}
)
