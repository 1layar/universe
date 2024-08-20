package dto

type PermissionDto struct {
	Role        string
	UserId      int
	Permissions []string
	Subject     string
	SubjectID   int
}

type PermissionCheckDTO struct {
	UserId     int
	Role       string
	Permission string
	Subject    string
	SubjectID  int
}
