package postgres

import (
	"context"
	"fmt"

	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgAuth "github.com/maruki00/deligo/pkg/auth"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type AuthzRepo struct {
	db    *pkgPostgres.PGHandler
	authz *pkgAuth.Authz
}

func NewAuthzRepo(db *pkgPostgres.PGHandler, authz *pkgAuth.Authz) *AuthzRepo {
	return &AuthzRepo{
		db:    db,
		authz: authz,
	}
}

func (_this *AuthzRepo) SavePolicy(ctx context.Context, policy *model.Policy) error {
	mapped := policy.Map2Casbin()

	err := _this.authz.AddPolicy(ctx, mapped.V0, mapped.V1, mapped.V2, *mapped.V3, *mapped.V4, *mapped.V5)
	if err != nil {
		return fmt.Errorf("error saving new policy: %v\n", err.Error())
	}
	return nil
}

func (_this *AuthzRepo) SaveGoupPolicy(ctx context.Context, group *model.GroupPolicy) error {
	mapped := group.Map2Casbin()
	return _this.authz.AddGroupingPolicy(ctx, mapped.V0, mapped.V1, mapped.V2)
}

func (_this *AuthzRepo) Check(ctx context.Context, sub, obj, act string) (bool, error) {
	return _this.authz.Check(ctx, sub, obj, act)
}
