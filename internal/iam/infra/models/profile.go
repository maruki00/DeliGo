package models

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	"time"
)

type Profile struct {
	ID        shared_valueobject.ID
	UserID    shared_valueobject.ID
	FullName  string
	Avatar    string
	Bio       string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
