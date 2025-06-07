package model

import shared_model "deligo/internal/shared/model"

type Menu struct {
	shared_model.BaseModel
	Label string `json:"label"`
}

func (_this *Menu) GetId() int {
	return _this.ID
}
func (_this *Menu) GetLAbel() string {
	return _this.Label
}
