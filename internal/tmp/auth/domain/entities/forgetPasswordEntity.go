package entities

type ForgetPasswordEntity interface {
	GetId() int
	GetEmail() string
	GetPin() int
	GetToken() string
}
