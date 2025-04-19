package models

import (
	"time"
)

type Group struct {
	ID string

	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
