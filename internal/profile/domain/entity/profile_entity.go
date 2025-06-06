package entity

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"time"
)

type ProfileEntity interface {
	GetID() shared_valueobject.ID
	GetUserID() shared_valueobject.ID
	GetFullName() string
	GetAvatar() string
	GetBio() string
	GetDeletedAt() *time.Time
	GetCreatedAt() *time.Time
	GetUpdatedAt() *time.Time
}
