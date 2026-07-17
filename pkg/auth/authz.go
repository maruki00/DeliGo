package pkgAuth

import (
	"context"
	"fmt"

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
	config string,
) (*Authz, error) {
	authz := Authz{}
	adapter, err := gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		return nil, err
	}
	engine, err := casbin.NewEnforcer(config, authz.adapter)
	if err != nil {
		return nil, err
	}
	if err := engine.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("Failed to load policies: %v\n", err.Error())
	}
	return &Authz{
		adapter: adapter,
		engine:  engine,
	}, nil
}

func (e *Authz) GetEngine() *casbin.Enforcer {
	return e.engine
}

func (e *Authz) AddPolicy(
	ctx context.Context,
	values ...string,
) error {
	_, err := e.engine.AddPolicy(values)
	if err != nil {
		return err
	}
	return e.engine.SavePolicy()
}

func (e *Authz) RmPolicy(
	ctx context.Context,
	values ...string,
) error {
	_, err := e.engine.RemovePolicy(values)
	if err != nil {
		return err
	}
	return e.engine.SavePolicy()
}

func (e *Authz) AddGroupingPolicy(
	ctx context.Context,
	values ...string,
) error {
	_, err := e.engine.AddGroupingPolicy(values)
	if err != nil {
		return err
	}
	return e.engine.SavePolicy()
}

func (e *Authz) RmGroupingPolicy(
	ctx context.Context,
	values ...string,
) error {
	_, err := e.engine.RemoveGroupingPolicy(values)
	if err != nil {
		return err
	}
	return e.engine.SavePolicy()
}

func (e *Authz) Check(
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
