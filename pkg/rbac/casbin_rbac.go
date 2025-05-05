package rbac

import (
	"github.com/casbin/casbin/v2"
)

var _ RBACManager = (*casbinRBAC)(nil)

type casbinRBAC struct {
	enforcer *casbin.Enforcer
}

func NewRBACManager(enforcer *casbin.Enforcer) RBACManager {
	return &casbinRBAC{enforcer: enforcer}
}

func (r *casbinRBAC) AddRoleForUser(userID, role string) error {
	_, err := r.enforcer.AddRoleForUser(userID, role)
	return err
}

func (r *casbinRBAC) DeleteRoleForUser(userID, role string) error {
	_, err := r.enforcer.DeleteRoleForUser(userID, role)
	return err
}

func (r *casbinRBAC) GetRolesForUser(userID string) ([]string, error) {
	return r.enforcer.GetRolesForUser(userID)
}

func (r *casbinRBAC) AddPermissionForRole(role, obj, act string) error {
	_, err := r.enforcer.AddPermissionForUser(role, obj, act)
	return err
}

func (r *casbinRBAC) DeletePermissionForRole(role, obj, act string) error {
	_, err := r.enforcer.DeletePermissionForUser(role, obj, act)
	return err
}

func (r *casbinRBAC) GetPermissionsForRole(role string) ([][2]string, error) {
	perms, err := r.enforcer.GetPermissionsForUser(role)
	if err != nil {
		return nil, err
	}

	var result [][2]string
	for _, p := range perms {
		if len(p) >= 3 {
			result = append(result, [2]string{p[1], p[2]})
		}
	}
	return result, nil
}

func (r *casbinRBAC) Enforce(userID, obj, act string) (bool, error) {
	return r.enforcer.Enforce(userID, obj, act)
}
