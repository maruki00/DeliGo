package models

import (
	"time"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Id        string    `json: "id"`
	Name      string    `json: "name"`
	OpenAt    time.Time `json: "open_at"`
	ClsoeAt   time.Time `json: "close_at"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func (_this *Shop) GetId() string {
	return _this.Id
}

func (_this *Shop) GetName() string {
	return _this.Name
}

func (_this *Shop) GetOpenAt() time.Time {
	return _this.OpenAt
}

func (_this *Shop) GetClsoeAt() time.Time {
	return _this.ClsoeAt
}

func (_this *Shop) GetCreatedAt() time.Time {
	return _this.CreatedAt
}

func (_this *Shop) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}
