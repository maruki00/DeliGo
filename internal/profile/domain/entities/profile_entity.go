package entities

import "time"

type ProfileEntity interface {
	SetID(ID string)
	SetUserID(UserID string)
	SetFullName(FullName string)
	SetAvatar(Avatar string)
	SetBio(Bio string)
	SetDeletedAt(DeletedAt *time.Time)
	SetCreatedAt(CreatedAt *time.Time)
	SetUpdatedAt(UpdatedAt *time.Time)
	GetID() string
	GetUserID() string
	GetFullName() string
	GetAvatar() string
	GetBio() string
	GetDeletedAt() *time.Time
	GetCreatedAt() *time.Time
	GetUpdatedAt() *time.Time
}
