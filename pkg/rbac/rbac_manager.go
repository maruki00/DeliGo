package rbac

type RBACManager interface {
	AddRoleForUser(userID, role string) error
	DeleteRoleForUser(userID, role string) error
	GetRolesForUser(userID string) ([]string, error)
	AddPermissionForRole(role string, obj, act string) error
	DeletePermissionForRole(role string, obj, act string) error
	GetPermissionsForRole(role string) ([][2]string, error)
	Enforce(userID, obj, act string) (bool, error)
}
