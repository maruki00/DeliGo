package PkgAuth

import (
	"context"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type Authz struct {
	adapter *gormadapter.Adapter
	engine  *casbin.Enforcer
}

func NewEnforcer(
	db *pkgPostgres.PGHandler,
	cfg string,
) *Authz {
	return &Authz{
		engine: engine,
	}
}

func (authz *Authz) GetEngine() *casbin.Enforcer {
	return e.engine
}

func (authz *Authz) AddPolicy(
	ctx context.Context,
	values ...string,
) error {

	_, err := e.engine.AddPolicy(values)

	if err != nil {
		return err
	}

	return e.engine.SavePolicy()
}

func (authz *Authz) AddGroupingPolicy(
	ctx context.Context,
	values ...string,
) error {

	_, err := e.engine.AddGroupingPolicy(values)

	if err != nil {
		return err
	}

	return e.engine.SavePolicy()
}

func (authz *Authz) Check(
	ctx context.Context,
	sub string,
	obj string,
	act string,
) (bool, error) {

	return e.engine.Enforce(
		sub,
		obj,
		act,
	)
}
