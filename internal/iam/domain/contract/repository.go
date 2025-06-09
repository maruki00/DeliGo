package contracts

import (
	"context"
	"deligo/internal/iam/domain/entity"
	"deligo/internal/iam/infra/model"
	sharedvo "deligo/internal/shared/valueobject"
	"deligo/pkg/pagination"
)

type IUserRepository interface {
	Save(context.Context, entity.UserEntity) error
	Delete(context.Context, sharedvo.ID) error
	Update(context.Context, sharedvo.ID, map[string]interface{}) error
	FindByID(context.Context, sharedvo.ID) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	FindByUsername(context.Context, string) (*model.User, error)
	ListByTenant(context.Context, sharedvo.ID, pagination.Pagination) ([]*model.User, error)
	AffectRole(context.Context, string) error
}

type IRoleRepositoty interface {
	Save(context.Context, entity.RoleEntity) error
	FindByID(context.Context, sharedvo.ID) (entity.RoleEntity, error)
	FindByName(context.Context, string) (entity.RoleEntity, error)
	List(context.Context, pagination.Pagination) ([]*model.Role, error)
	AffectPolicy(context.Context, string, string, string) error
	Delete(context.Context, sharedvo.ID) error
}

type IPolicyRepository interface {
	Save(context.Context, *model.Policy) error
	Delete(context.Context, string) error
	FindByID(context.Context, string) (*model.Policy, error)
	FindByName(context.Context, string) (*model.Policy, error)
	AffectPermission(context.Context, string, string, string) error
	ListForTenant(context.Context, string, pagination.Pagination) ([]*model.Policy, error)
}

type IPermessionRepository interface {
	Save(context.Context, entity.PermissionEntity) error
	FindByID(context.Context, string) (*model.Permission, error)
	FindByPolicyID(context.Context, string, pagination.Pagination) ([]*model.Permission, error)
	Delete(context.Context, string) error
}
