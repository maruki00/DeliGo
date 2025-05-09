package entities

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	"time"
)

type PolicyEntity interface {
	SetID(_ID valueobjects.ID)
	SetName(_Name string)
	GetID() valueobjects.ID
	GetName() string
	GetGroupID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetPermissions() []*models.Permission
}
