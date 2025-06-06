package entity

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
)

type PermissionEntity interface {
	GetID() valueobjects.ID
	GetName() string
	GetAction() string
	GetDescription() string
}
