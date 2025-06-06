package model

import (
	shared_model "deligo/internal/shared/infra/model"
)

type Menu struct {
	shared_model.BaseModel
	Label string `json: "label"`
}

func (obj *Menu) GetId() int {
	return obj.Id
}
func (obj *Menu) GetLAbel() string {
	return obj.LAbel
}
