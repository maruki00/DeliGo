package models

import (
	shared_models "deligo/internal/shared/infra/models"
	"time"
)

type Shop struct {
	shared_models.BaseModel
	Id      string    `json: "id"`
	Name    string    `json: "name"`
	OpenAt  time.Time `json: "open_at"`
	ClsoeAt time.Time `json: "close_at"`
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
