package entities

import "time"

type ProfileEntity interface {
	GetID() string
	GetUserID() string
	GetFullName() string
	GetAvatar() string
	GetBio() string
	GetDeletedAt() *time.Time
	GetCreatedAt() *time.Time
	GetUpdatedAt() *time.Time
}
