package entity

import valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"

type RoleEntity interface {
	GetID() valueobjects.ID
	GetName() string
	GetDescription() string
}
