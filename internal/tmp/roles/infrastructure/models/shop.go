package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Id        int       `json: "id"`
	Name      string    `json: "name"`
	OpenAt    time.Time `json: "open_at"`
	ClsoeAt   time.Time `json: "close_at"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func (obj *Shop) GetId() int {
	return obj.Id
}

func (obj *Shop) GetName() string {
	return obj.Name
}

func (obj *Shop) GetOpenAt() time.Time {
	return obj.OpenAt
}

func (obj *Shop) GetClsoeAt() time.Time {
	return obj.ClsoeAt
}

func (obj *Shop) GetCreatedAt() time.Time {
	return obj.CreatedAt
}

func (obj *Shop) GetUpdatedAt() time.Time {
	return obj.UpdatedAt
}
