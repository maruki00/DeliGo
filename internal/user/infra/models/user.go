package models

import "time"

type User struct {
	ID        string
	Email     string
	Password  string
	Role      string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
