package models

import "time"

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

func (_this *Profile) GetID() string {
	return _this.ID
}
func (_this *Profile) GetUserID() string {
	return _this.UserID
}
func (_this *Profile) GetFullName() string {
	return _this.FullName
}
func (_this *Profile) GetAvatar() string {
	return _this.Avatar
}
func (_this *Profile) GetBio() string {
	return _this.Bio
}
func (_this *Profile) GetDeletedAt() *time.Time {
	return _this.DeletedAt
}
func (_this *Profile) GetCreatedAt() *time.Time {
	return _this.CreatedAt
}
func (_this *Profile) GetUpdatedAt() *time.Time {
	return _this.UpdatedAt
}
