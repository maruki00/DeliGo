package model

import "time"

type AUTHN struct {
	LoginID     string
	Password    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (_this *AUTHN) GetLoginID() string {
	return _this.LoginID
}
func (_this *AUTHN) GetPassword() string {
	return _this.Password
}
func (_this *AUTHN) GetDescription() string {
	return _this.Description
}

func (_this *AUTHN) SetLoginID(LoginID string) {
	_this.LoginID = LoginID
}
func (_this *AUTHN) SetPassword(Password string) {
	_this.Password = Password
}
func (_this *AUTHN) SetDescription(Description string) {
	_this.Description = Description
}

func (_this *AUTHN) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *AUTHN) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

func (_this *AUTHN) SetCreatedAt(CreatedAt time.Time) {
	_this.CreatedAt = CreatedAt
}
func (_this *AUTHN) SetUpdatedAt(UpdatedAt time.Time) {
	_this.UpdatedAt = UpdatedAt
}
