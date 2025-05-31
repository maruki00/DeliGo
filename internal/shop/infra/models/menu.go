package models

import (
	shared_models "deligo/internal/shared/infra/models"
)

type Menu struct {
	shared_models.BaseModel
	Label string `json: "label"`
}

func (obj *Menu) GetId() int {
	return obj.Id
}
func (obj *Menu) GetLAbel() string {
	return obj.LAbel
}
