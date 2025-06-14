package entity

import (
	"time"

	shared_valueobject "github.com/maruki00/deligo/internal/shared/valueobject"
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
