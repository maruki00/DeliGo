package entities

import valueobjects "deligo/internal/iam/domain/valueobject"

type RoleEntity interface {
	GetID() valueobjects.ID
	GetName() string
}
