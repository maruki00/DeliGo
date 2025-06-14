package entity

import (
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
)

type PermissionEntity interface {
	GetID() valueobjects.ID
	GetName() string
	GetAction() string
	GetDescription() string
}
