package contracts

import (
	"context"
	"deligo/internal/iam/domain/entity"
	"deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/model"
	"deligo/pkg/pagination"
)

type IUserRepository interface {
	Save(context.Context, entity.UserEntity) error
	Delete(context.Context, valueobject.ID) error
	Update(context.Context, valueobject.ID, map[string]interface{}) error
	FindByID(context.Context, valueobject.ID) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	FindByUsername(context.Context, string) (*model.User, error)
	ListByTenant(context.Context, valueobject.ID, pagination.Pagination) ([]*model.User, error)
	AffectRole(context.Context, string) error
}
