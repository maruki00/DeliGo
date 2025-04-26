package models

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	UserId    int    `json:"user_id"`
	UserType  string `json:"user_type"`
	UserLevel int    `json:"user_level"`
	ExpiresAt string `json:"expires_at"`
}

func (obj *Auth) SetId(id int) {
	obj.Id = id
}
func (obj *Auth) SetEmail(email string) {
	obj.Email = email
}
func (obj *Auth) SetToken(token string) {
	obj.Token = token
}
func (obj *Auth) SetUserId(user_id int) {
	obj.UserId = user_id
}
func (obj *Auth) SetUserType(user_type string) {
	obj.UserType = user_type
}
func (obj *Auth) SetUserLevel(user_level int) {
	obj.UserLevel = user_level
}

func (obj *Auth) GetId() int {
	return obj.Id
}
func (obj *Auth) GetEmail() string {
	return obj.Email
}
func (obj *Auth) GetToken() string {
	return obj.Token
}
func (obj *Auth) GetUserId() int {
	return obj.UserId
}
func (obj *Auth) GetUserType() string {
	return obj.UserType
}
func (obj *Auth) GetUserLevel() int {
	return obj.UserLevel
}
