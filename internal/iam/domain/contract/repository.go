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
	Save(ctx context.Context, entity *model.Policy) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*model.Policy, error)
	FindByName(ctx context.Context, name string) (*model.Policy, error)
	AffectPermission(ctx context.Context, id, policy_id string, permission_id string) error
	ListForTenant(ctx context.Context, id string, pagination pagination.Pagination) ([]*model.Policy, error)
}

type IPermessionRepository interface {
	Save(ctx context.Context, permission entity.PermissionEntity) error
	FindByID(ctx context.Context, id string) (*model.Permission, error)
	FindByPolicyID(ctx context.Context, policyID string, pagination model.Pagination) ([]*model.Permission, error)
	Delete(ctx context.Context, id string) error
}
