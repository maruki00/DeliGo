package contracts

import (
	"context"
	"deligo/internal/iam/domain/entities"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
)

type Pagination struct{}

type IUserRepository interface {
	Save(context.Context, entities.UserEntity) error
	Delete(context.Context, valueobjects.ID) error
	Update(context.Context, valueobjects.ID, map[string]interface{}) error
	FindByID(context.Context, valueobjects.ID) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	ListByTenant(context.Context, valueobjects.ID, shared_models.Pagination) ([]*models.User, error)
}

type IPolicyRepository interface {
	Save(context.Context, *models.Policy) error
	Delete(context.Context, string) error
	FindByID(context.Context, string) (*models.Policy, error)
	FindByName(context.Context, string) (*models.Policy, error)
	ListForTenant(context.Context, string) ([]*models.Policy, error)
}

type IRoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id string) (*models.Role, error)
	GetByName(ctx context.Context, name string) (*models.Role, error)
	List(ctx context.Context) ([]*models.Role, error)
	Delete(ctx context.Context, id string) error
}

type IPermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) error
	GetByID(ctx context.Context, id string) (*models.Permission, error)
	GetByName(ctx context.Context, name string) (*models.Permission, error)
	List(ctx context.Context) ([]*models.Permission, error)
	Delete(ctx context.Context, id string) error
}
type IGroupRepository interface {
	AssignUserToGroup(ctx context.Context, userID, groupID string) error
	RemoveUserFromGroup(ctx context.Context, userID, groupID string) error
	GetUserGroups(ctx context.Context, userID string) ([]*models.Group, error)
}
