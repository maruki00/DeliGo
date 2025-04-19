package contracts

import (
	"context"
	"deligo/internal/iam/infra/models"
)

type Pagination struct{}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	ListByTenant(ctx context.Context, tenantID string, pagination Pagination) ([]*models.User, error)
}

type PolicyRepository interface {
	FindByID(ctx context.Context, id string) (*models.Policy, error)
	FindByName(ctx context.Context, name string) (*models.Policy, error)
	Save(ctx context.Context, policy *models.Policy) error
	Delete(ctx context.Context, id string) error
	ListForTenant(ctx context.Context, tenantID string) ([]*models.Policy, error)
}

type PermissionRepository interface {
	FindByID(ctx context.Context, id string) (*models.Permission, error)
	FindByPolicyID(ctx context.Context, policyID string) ([]*models.Permission, error)
	Save(ctx context.Context, permission *models.Permission) error
	Delete(ctx context.Context, id string) error
}

type RoleRepository interface {
	AssignUserToRole(ctx context.Context, userID, roleID string) error
	RemoveUserFromRole(ctx context.Context, userID, roleID string) error
	GetUserRoles(ctx context.Context, userID string) ([]*models.Role, error)
}
