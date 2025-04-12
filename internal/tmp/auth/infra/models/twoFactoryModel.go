package models

import "gorm.io/gorm"

type TwoFactoryPin struct {
	gorm.Model
	Pin   int    `json:"pin"`
	Email string `json:"email"`
}

func (o *TwoFactoryPin) GetPin() int {
	return o.Pin
}

func (o *TwoFactoryPin) GetEmail() string {
	return o.Email
}
