package models

import (
	"time"
)

type Profile struct {
	ID        string
	UserID    string
	FullName  string
	Avatar    string
	Bio       string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
