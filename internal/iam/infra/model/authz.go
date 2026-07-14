package model

import (
	"time"
)

type AUTHZ struct {
	ID          string
	Name        string
	Action      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (_this *AUTHZ) GetID() string {
	return _this.ID
}
func (_this *AUTHZ) GetName() string {
	return _this.Name
}
func (_this *AUTHZ) GetAction() string {
	return _this.Action
}
func (_this *AUTHZ) GetDescription() string {
	return _this.Description
}

func (_this *AUTHZ) SetID(ID string) {
	_this.ID = ID
}
func (_this *AUTHZ) SetName(Name string) {
	_this.Name = Name
}
func (_this *AUTHZ) SetAction(Action string) {
	_this.Action = Action
}
func (_this *AUTHZ) SetDescription(Description string) {
	_this.Description = Description
}
func (_this *AUTHZ) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *AUTHZ) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

func (_this *AUTHZ) SetCreatedAt(CreatedAt time.Time) {
	_this.CreatedAt = CreatedAt
}
func (_this *AUTHZ) SetUpdatedAt(UpdatedAt time.Time) {
	_this.UpdatedAt = UpdatedAt
}
