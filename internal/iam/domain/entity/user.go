package entity

import (
	"time"
)

type UserRole string
type UserStatus string

const (
	RoleAdmin           UserRole = "admin"
	RoleCustomer        UserRole = "customer"
	RoleRestaurantOwner UserRole = "restaurant_owner"
	RoleCourier         UserRole = "courier"

	StatusCreated   UserStatus = "created"
	StatusActive    UserStatus = "active"
	StatusSuspended UserStatus = "suspended"
	StatusBanned    UserStatus = "banned"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Phone        string
	FirstName    string
	LastName     string
	Role         UserRole
	Status       UserStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) Activate() {
	u.Status = StatusActive
	u.UpdatedAt = time.Now()
}

func (u *User) Ban() {
	u.Status = StatusBanned
	u.UpdatedAt = time.Now()
}
