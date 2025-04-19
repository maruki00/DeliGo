package contracts

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	ListByTenant(ctx context.Context, tenantID string, pagination Pagination) ([]*User, error)
}

type PolicyRepository interface {
	FindByID(ctx context.Context, id string) (*Policy, error)
	FindByName(ctx context.Context, name string) (*Policy, error)
	Save(ctx context.Context, policy *Policy) error)
	Delete(ctx context.Context, id string) error
	ListForTenant(ctx context.Context, tenantID string) ([]*Policy, error)
}

type PermissionRepository interface {
	FindByID(ctx context.Context, id string) (*Permission, error)
	FindByPolicyID(ctx context.Context, policyID string) ([]*Permission, error)
	Save(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
}

type RoleRepository interface {
	AssignUserToRole(ctx context.Context, userID, roleID string) error
	RemoveUserFromRole(ctx context.Context, userID, roleID string) error
	GetUserRoles(ctx context.Context, userID string) ([]*Role, error)
}