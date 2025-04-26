package models

import "gorm.io/gorm"

type ResetPassword struct {
	gorm.Model
	Id    int    `json:"id"`
	Email string `json:"email"`
	Pin   int    `json:"pin"`
	Token string `json:"token"`
	// CreatedAt string
	// UpdatedAt string
}

func (obj *ResetPassword) GetId() int {
	return obj.Id
}
func (obj *ResetPassword) GetEmail() string {
	return obj.Email
}
func (obj *ResetPassword) GetToken() string {
	return obj.Token
}
