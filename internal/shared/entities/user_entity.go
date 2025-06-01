package shared_entities

type UserEntity interface {
	GetID() int
	GetUsername() string
	GetFullname() string
	GetEmail() string
	GetAddress() string
	GetPassword() string
	GetUserType() string
	GetUserLevel() int
	GetIsOnline() int
	GetIsLocked() int
	GetLastSeen() string
	GetCreatedAt() string
	GetUpdatedAt() string
}
