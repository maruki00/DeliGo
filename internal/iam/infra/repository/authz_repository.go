package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type AuthzRepository struct {
	db       *pkgPostgres.PGHandler
	adapter  *gormadapter.Adapter
	enforcer *casbin.Enforcer
}

func NewAuthzRepository(db *pkgPostgres.PGHandler, config string) (*AuthzRepository, error) {
	authz := AuthzRepository{
		db: db,
	}
	var err error
	authz.adapter, err = gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		panic(err)
	}

	authz.enforcer, err = casbin.NewEnforcer(config, authz.adapter)
	if err != nil {
		panic(err)
	}

	if err := authz.enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("Failed to load policies: %v\n", err.Error())
	}

	authz.enforcer.ClearPolicy()
	return &authz, nil
}

func (_this *AuthzRepository) SavePolicy(ctx context.Context, policy model.Policy) error {
	mapped := policy.Map2Casbin()

	ok, err := _this.enforcer.AddPolicy(mapped.V0, mapped.V1, mapped.V2, mapped.V3, mapped.V4, mapped.V5)
	if err != nil {
		return fmt.Errorf("error saving new policy: %v\n", err.Error())
	}
	if !ok {
		return errors.New("could not add new policy")
	}
	if err := _this.enforcer.SavePolicy(); err != nil {
		return fmt.Errorf("Failed to save policies: %v", err)
	}
	return nil
}

func (_this *AuthzRepository) SaveGoupPolicy(ctx context.Context, group model.GroupPolicy) error {
	mapped := group.Map2Casbin()
	ok, err := _this.enforcer.AddGroupingPolicy(mapped.V0, mapped.V1, mapped.V2)
	if err != nil {
		return fmt.Errorf("error saving new group policy: %v\n", err.Error())
	}

	if !ok {
		return errors.New("could not add new group policy")
	}

	if err := _this.enforcer.SavePolicy(); err != nil {
		return fmt.Errorf("Failed to save policies: %v", err)
	}
	return nil
}

func (_this *AuthzRepository) Check(ctx context.Context, sub, obj, act string) (bool, error) {
	allowed, err := _this.enforcer.Enforce(sub, obj, act)
	if err != nil {

		return false, fmt.Errorf("Error validating rule: %v\n", err)
	}
	return allowed, nil
}
