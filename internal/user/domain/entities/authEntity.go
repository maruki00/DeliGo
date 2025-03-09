package entities

type AuthEntity interface {
	SetId(int)
	SetEmail(string)
	SetToken(string)
	SetUserId(int)
	SetUserType(string)
	SetUserLevel(string)
	GetId() int
	GetEmail() string
	GetToken() string
	GetUserId() int
	GetUserType() string
	GetUserLevel() string
}
