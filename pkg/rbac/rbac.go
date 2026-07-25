package rbac

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"gorm.io/gorm"
)

func NewRBAC(db *gorm.DB, rbac_model string) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	//"rbac_model.conf"
	enforcer := casbin.NewEnforcer(rbac_model, adapter)

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}
