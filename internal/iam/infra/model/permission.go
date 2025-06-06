package model

import (
	valueobject "deligo/internal/iam/domain/valueobject"
)

type Permission struct {
	ID          valueobject.ID
	Name        string
	Action      string
	Description string
}

func (_this *Permission) GetID() valueobject.ID {
	return _this.ID
}
func (_this *Permission) GetName() string {
	return _this.Name
}
func (_this *Permission) GetAction() string {
	return _this.Action
}
func (_this *Permission) GetDescription() string {
	return _this.Description
}
