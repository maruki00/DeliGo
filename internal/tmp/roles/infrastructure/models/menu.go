package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Id        int       `json: "id"`
	LAbel     string    `json: "label"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func (obj *Menu) GetId() int {
	return obj.Id
}
func (obj *Menu) GetLAbel() string {
	return obj.LAbel
}
