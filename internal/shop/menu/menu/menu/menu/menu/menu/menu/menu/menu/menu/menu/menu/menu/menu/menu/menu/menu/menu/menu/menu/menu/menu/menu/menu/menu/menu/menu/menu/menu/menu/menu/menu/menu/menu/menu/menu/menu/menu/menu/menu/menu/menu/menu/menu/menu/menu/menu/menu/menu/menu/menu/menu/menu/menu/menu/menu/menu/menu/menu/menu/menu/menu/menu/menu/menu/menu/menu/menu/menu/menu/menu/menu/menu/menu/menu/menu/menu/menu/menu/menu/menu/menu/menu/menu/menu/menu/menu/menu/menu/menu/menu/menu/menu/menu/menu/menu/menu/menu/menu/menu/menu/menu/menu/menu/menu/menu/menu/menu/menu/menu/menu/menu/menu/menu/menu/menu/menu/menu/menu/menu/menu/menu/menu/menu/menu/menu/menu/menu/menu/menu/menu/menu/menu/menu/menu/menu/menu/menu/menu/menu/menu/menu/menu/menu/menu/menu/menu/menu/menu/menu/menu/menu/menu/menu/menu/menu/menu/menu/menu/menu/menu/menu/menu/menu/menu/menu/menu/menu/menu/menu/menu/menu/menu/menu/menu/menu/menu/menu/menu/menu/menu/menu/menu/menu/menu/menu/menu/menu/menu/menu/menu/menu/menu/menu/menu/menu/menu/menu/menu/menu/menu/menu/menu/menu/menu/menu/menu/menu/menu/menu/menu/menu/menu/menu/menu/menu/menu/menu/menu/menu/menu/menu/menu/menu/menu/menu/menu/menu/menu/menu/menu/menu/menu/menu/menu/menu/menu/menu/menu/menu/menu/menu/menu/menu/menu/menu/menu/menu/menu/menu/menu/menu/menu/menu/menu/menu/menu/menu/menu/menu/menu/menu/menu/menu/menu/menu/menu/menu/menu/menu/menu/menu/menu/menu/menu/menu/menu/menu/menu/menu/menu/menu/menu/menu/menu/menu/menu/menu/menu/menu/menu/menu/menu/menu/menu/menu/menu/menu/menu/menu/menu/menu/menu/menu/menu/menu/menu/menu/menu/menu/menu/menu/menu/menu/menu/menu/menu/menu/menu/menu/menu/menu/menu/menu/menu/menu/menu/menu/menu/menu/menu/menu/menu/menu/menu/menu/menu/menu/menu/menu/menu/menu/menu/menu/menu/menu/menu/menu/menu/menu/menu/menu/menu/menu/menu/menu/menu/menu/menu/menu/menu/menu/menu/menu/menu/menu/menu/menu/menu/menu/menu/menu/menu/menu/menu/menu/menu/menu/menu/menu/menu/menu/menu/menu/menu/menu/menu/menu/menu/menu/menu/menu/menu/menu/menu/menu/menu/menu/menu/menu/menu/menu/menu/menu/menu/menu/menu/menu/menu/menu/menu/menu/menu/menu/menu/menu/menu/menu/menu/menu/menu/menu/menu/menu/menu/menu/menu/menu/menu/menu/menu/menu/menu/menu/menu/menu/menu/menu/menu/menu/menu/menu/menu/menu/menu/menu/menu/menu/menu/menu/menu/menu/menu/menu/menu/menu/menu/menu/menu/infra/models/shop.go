package models

import (
	shared_models "deligo/internal/shared/infra/models"
	"time"
)

type Shop struct {
	shared_models.BaseModel
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
