package models

import (
	"delivery/internal/user/domain/entities"
	"time"
)

var _ entities.UserEntity = (*User)(nil)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *User) SetID(ID string) {
	u.ID = ID
}
func (u *User) SetEmail(Email string) {
	u.Email = Email
}
func (u *User) SetPassword(Password string) {
	u.Password = Password
}
func (u *User) SetRole(Role string) {
	u.Role = Role
}
func (u *User) SetDeletedAt(DeletedAt *time.Time) {
	u.DeletedAt = DeletedAt
}
func (u *User) SetCreatedAt(CreatedAt *time.Time) {
	u.CreatedAt = CreatedAt
}
func (u *User) SetUpdatedAt(UpdatedAt *time.Time) {
	u.UpdatedAt = UpdatedAt
}

func (u *User) GetID() string {
	return u.ID
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetPassword() string {
	return u.Password
}
func (u *User) GetRole() string {
	return u.Role
}
func (u *User) GetDeletedAt() *time.Time {
	return u.DeletedAt
}
func (u *User) GetCreatedAt() *time.Time {
	return u.CreatedAt
}
func (u *User) GetUpdatedAt() *time.Time {
	return u.UpdatedAt
}
