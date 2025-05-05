package entities

import "time"

type PermissionEntity interface {
	SetID(ID string)
	SetEmail(Email string)
	SetPassword(Password string)
	SetRole(Role string)
	SetDeletedAt(DeletedAt *time.Time)
	SetCreatedAt(CreatedAt *time.Time)
	SetUpdatedAt(UpdatedAt *time.Time)
	GetID() string
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetDeletedAt() *time.Time
	GetCreatedAt() *time.Time
	GetUpdatedAt() *time.Time
}
