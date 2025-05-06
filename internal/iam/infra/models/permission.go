package models

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
)

type Permission struct {
	ID          valueobjects.ID
	Name        string
	Action      string
	Description string
}

func (_this *Permission) GetID() valueobjects.ID {
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
