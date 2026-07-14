package casbin

import (
	"context"

	"github.com/casbin/casbin/v3"
)

type Enforcer struct {
	engine *casbin.Enforcer
}

func New(
	engine *casbin.Enforcer,
) *Enforcer {
	return &Enforcer{
		engine: engine,
	}
}

func (e *Enforcer) AddPolicy(
	ctx context.Context,
	values ...interface{},
) error {

	_, err := e.engine.AddPolicy(values...)

	if err != nil {
		return err
	}

	return e.engine.SavePolicy()
}

func (e *Enforcer) AddGroupingPolicy(
	ctx context.Context,
	values ...interface{},
) error {

	_, err := e.engine.AddGroupingPolicy(values...)

	if err != nil {
		return err
	}

	return e.engine.SavePolicy()
}

func (e *Enforcer) Check(
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
