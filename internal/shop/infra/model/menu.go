package model

import shared_model "deligo/internal/shared/model"

type Menu struct {
	shared_model.BaseModel
	Label string `json: "label"`
}

func (obj *Menu) GetId() int {
	return obj.ID
}
func (obj *Menu) GetLAbel() string {
	return obj.Label
}
