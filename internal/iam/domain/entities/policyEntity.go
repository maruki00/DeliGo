package entities

import (
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	"time"

	"gorm.io/gorm"
)

type PolicyEntity interface {
	SetID(_ID valueobjects.ID)
	SetName(_Name string)
	SetGroupID(_GroupID string)

	GetID() valueobjects.ID
	GetName() string
	GetGroupID() string
	GetDeletedAt() gorm.DeletedAt
	GetUpdatedAt() time.Time
	GetPermissions() []*models.Permission
	GetGroups() []*models.Group
}
