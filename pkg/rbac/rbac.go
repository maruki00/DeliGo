package rbac

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"gorm.io/gorm"
)

func NewRBAC(db *gorm.DB, rbac_model string) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	//"rbac_model.conf"
	enforcer, err := casbin.NewEnforcer(rbac_model, adapter)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}
