package entity

import (
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	"github.com/maruki00/deligo/internal/iam/infra/model"
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
	GetPermissions() []*model.Permission
}
