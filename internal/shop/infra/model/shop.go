package model

import (
	shared_model "deligo/internal/shared/infra/model"
	"time"
)

type Shop struct {
	shared_model.BaseModel
	Name    string    `json:"name,omitempty"`
	OpenAt  time.Time `json:"open_at,omitempty"`
	CloseAt time.Time `json:"clsoe_at,omitempty"`
}

func (_this *Shop) GetName() string {
	return _this.Name
}

func (_this *Shop) GetOpenAt() time.Time {
	return _this.OpenAt
}

func (_this *Shop) GetCloseAt() time.Time {
	return _this.CloseAt
}
