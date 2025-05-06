package models

import valueobjects "deligo/internal/iam/domain/valueobject"

type Role struct {
	ID          valueobjects.ID
	Name        string
	Description string
}

func (_this *Role) GetID() valueobjects.ID {
	return _this.ID
}

func (_this *Role) GetName() string {
	return _this.Name
}

func (_this *Role) GetDescription() string {
	return _this.Description
}
