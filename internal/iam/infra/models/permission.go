package models

import (
	"time"
)

type Permission struct {
	ID string

	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
