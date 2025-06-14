package contracts

import (
	"context"
	"github.com/maruki00/deligo/internal/iam/domain/entity"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/valueobject"
	"github.com/maruki00/deligo/pkg/pagination"
)

type IUserRepository interface {
	Save(context.Context, entity.UserEntity) error
	Delete(context.Context, sharedvo.ID) error
	Update(context.Context, sharedvo.ID, map[string]interface{}) error
	FindByID(context.Context, sharedvo.ID) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	FindByUsername(context.Context, string) (*model.User, error)
	ListByTenant(context.Context, sharedvo.ID, pagination.Pagination) ([]*model.User, error)
	AffectRole(ctx context.Context, id sharedvo.ID, role_id string) error
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

type IPermissionRepository interface {
	Save(context.Context, entity.PermissionEntity) error
	FindByID(context.Context, string) (*model.Permission, error)
	FindByPolicyID(context.Context, string, pagination.Pagination) ([]*model.Permission, error)
	Delete(context.Context, string) error
}
