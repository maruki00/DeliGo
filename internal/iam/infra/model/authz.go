package model

import (
	"time"
)

type Authz struct {
	ID          string
	Name        string
	Action      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (_this *Authz) GetID() string {
	return _this.ID
}
func (_this *Authz) GetName() string {
	return _this.Name
}
func (_this *Authz) GetAction() string {
	return _this.Action
}
func (_this *Authz) GetDescription() string {
	return _this.Description
}

func (_this *Authz) SetID(ID string) {
	_this.ID = ID
}
func (_this *Authz) SetName(Name string) {
	_this.Name = Name
}
func (_this *Authz) SetAction(Action string) {
	_this.Action = Action
}
func (_this *Authz) SetDescription(Description string) {
	_this.Description = Description
}
func (_this *Authz) GetCreatedAt() time.Time {
	return _this.CreatedAt
}
func (_this *Authz) GetUpdatedAt() time.Time {
	return _this.UpdatedAt
}

func (_this *Authz) SetCreatedAt(CreatedAt time.Time) {
	_this.CreatedAt = CreatedAt
}
func (_this *Authz) SetUpdatedAt(UpdatedAt time.Time) {
	_this.UpdatedAt = UpdatedAt
}
