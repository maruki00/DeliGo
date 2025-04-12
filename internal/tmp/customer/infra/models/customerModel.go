package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Id int `json:"id"`
}
