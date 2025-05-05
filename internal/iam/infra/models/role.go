package models

import valueobjects "deligo/internal/iam/domain/valueobject"

type Role struct {
	ID   valueobjects.ID
	Name string
}

func (_this *Role) GetID() valueobjects.ID {
	return _this.ID
}

func (_this *Role) GetName() string {
	return _this.Name
}
