package PkgAuth

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v3"
)

type Authz struct {
	engine *casbin.Enforcer
}

func NewEnforcer(
	adapter interface{},
	cfg string,
) (*Authz, error) {
	engine, err := casbin.NewEnforcer(cfg, adapter)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Casbin enforcer: %v", err)
	}
	return &Authz{
		engine: engine,
	}, nil
}

func (authz *Authz) GetEngine() *casbin.Enforcer {
	return authz.engine
}

func (authz *Authz) AddPolicy(
	ctx context.Context,
	values ...string,
) error {

	_, err := authz.engine.AddPolicy(values)

	if err != nil {
		return err
	}

	return authz.engine.SavePolicy()
}

func (authz *Authz) AddGroupingPolicy(
	ctx context.Context,
	values ...string,
) error {

	_, err := authz.engine.AddGroupingPolicy(values)

	if err != nil {
		return err
	}

	return authz.engine.SavePolicy()
}

func (authz *Authz) Check(
	ctx context.Context,
	sub string,
	obj string,
	act string,
) (bool, error) {

	return authz.engine.Enforce(
		sub,
		obj,
		act,
	)
}
