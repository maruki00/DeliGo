package contract

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/infra/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserGormModel) error
	GetByID(ctx context.Context, id string) (*model.UserGormModel, error)
	GetByEmail(ctx context.Context, email string) (*model.UserGormModel, error)
	Update(ctx context.Context, user *model.UserGormModel) error
}

type IdentityRepository interface {
	AddPolicy(sub, obj, act string) (bool, error)
	AddGroupPolicy(sub, group string) (bool, error)
	RemovePolicy(sub, obj, act string) (bool, error)
	RemoveGroupPolicy(sub, group string) (bool, error)
	CheckPermission(sub, obj, act string) (bool, error)
}
