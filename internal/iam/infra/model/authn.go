package model

import "time"

type Authn struct {
	LoginID     string
	Password    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (_this *Authn) GetLoginID() string {
	return _this.LoginID
}
func (_this *Authn) GetPassword() string {
	return _this.Password
}
func (_this *Authn) GetDescription() string {
	return _this.Description
}

func (_this *Authn) SetLoginID(LoginID string) {
	_this.LoginID = LoginID
}
func (_this *Authn) SetPassword(Password string) {
	_this.Password = Password
}
func (_this *Authn) SetDescription(Description string) {
	_this.Description = Description
}

func (_this *Authn) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *Authn) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

func (_this *Authn) SetCreatedAt(CreatedAt time.Time) {
	_this.CreatedAt = CreatedAt
}
func (_this *Authn) SetUpdatedAt(UpdatedAt time.Time) {
	_this.UpdatedAt = UpdatedAt
}
