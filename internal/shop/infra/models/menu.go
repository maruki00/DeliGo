package models

import (
	"time"
)

type Menu struct {
	Id        string    `json: "id"`
	LAbel     string    `json: "label"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	DeletedAt DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (obj *Menu) GetId() int {
	return obj.Id
}
func (obj *Menu) GetLAbel() string {
	return obj.LAbel
}
