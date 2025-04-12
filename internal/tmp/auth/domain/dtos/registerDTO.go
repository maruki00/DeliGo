package dtos

type RegisterDTO struct {
	UserName  string
	FullName  string
	Email     string
	Address   string
	Password  string
	UserType  string
	UserLevel int
	IsOnline  int
	IsLocked  int
	LastSeen  string
}
