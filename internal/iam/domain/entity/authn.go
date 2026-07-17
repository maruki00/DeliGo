package entity

type AuthnEntity interface {
	GetLoginID() string
	GetPassword() string
	GetDescription() string

	SetLoginID(LoginID string)
	SetPassword(Password string)
	SetDescription(Description string)
}
