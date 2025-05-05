package contracts

import (
	"context"
	"deligo/internal/iam/domain/entities"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
)

type IUserRepository interface {
	Save(context.Context, entities.UserEntity) error
	Delete(context.Context, valueobjects.ID) error
	Update(context.Context, valueobjects.ID, map[string]interface{}) error
	FindByID(context.Context, valueobjects.ID) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	ListByTenant(context.Context, valueobjects.ID, shared_models.Pagination) ([]*models.User, error)
}

type IRoleRepository interface {
	Create(ctx context.Context, role entities.RoleEntity) error
	GetByID(ctx context.Context, id string) (entities.RoleEntity, error)
	GetByName(ctx context.Context, name string) (entities.RoleEntity, error)
	List(ctx context.Context) ([]entities.RoleEntity, error)
	Delete(ctx context.Context, id string) error
}

type IUserRoleRepository interface {
	AssignRole(ctx context.Context, userID, roleID string) error
	RemoveRole(ctx context.Context, userID, roleID string) error
	GetRolesByUserID(ctx context.Context, userID string) ([]entities.RoleEntity, error)
}

type IPolicyRepository interface {
	Save(context.Context, *models.Policy) error
	Delete(context.Context, string) error
	FindByID(context.Context, string) (*models.Policy, error)
	FindByName(context.Context, string) (*models.Policy, error)
	ListForTenant(context.Context, string) ([]*models.Policy, error)
}

type IPermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) error
	GetByID(ctx context.Context, id string) (*models.Permission, error)
	GetByName(ctx context.Context, name string) (*models.Permission, error)
	List(ctx context.Context) ([]*models.Permission, error)
	Delete(ctx context.Context, id string) error
}

type IRolePermissionRepository interface {
	AttachPermission(ctx context.Context, roleID, permissionID string) error
	DetachPermission(ctx context.Context, roleID, permissionID string) error
	GetPermissionsByRoleID(ctx context.Context, roleID string) ([]*models.Permission, error)
}

type IRBACManager interface {
	AddRoleForUser(userID, roleName string) error
	DeleteRoleForUser(userID, roleName string) error
	GetRolesForUser(userID string) ([]string, error)
	AddPermissionForRole(roleName string, permissions ...string) error
	DeletePermissionForRole(roleName string, permissions ...string) error
	GetPermissionsForRole(roleName string) ([]string, error)
	Enforce(userID, obj, act string) (bool, error)
}
