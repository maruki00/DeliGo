package model

import (
	shared_model "deligo/internal/shared/model"
	"time"
)

type Shop struct {
	shared_model.BaseModel
	Name    string    `json:"name,omitempty"`
	Status  bool      `json:"status,omitempty"`
	OpenAt  time.Time `json:"open_at,omitempty"`
	CloseAt time.Time `json:"close_at,omitempty"`
}

func (_this *Shop) GetName() string {
	return _this.Name
}

func (_this *Shop) GetStatus() bool {
	return _this.Status
}

func (_this *Shop) GetOpenAt() time.Time {
	return _this.OpenAt
}

func (_this *Shop) GetCloseAt() time.Time {
	return _this.CloseAt
}
