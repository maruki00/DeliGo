package postgres

import (
	"context"
	"fmt"

	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgAuth "github.com/maruki00/deligo/pkg/auth"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type IdentityRepo struct {
	authz *pkgAuth.Authz
}

func NewIdentityRepo(db *pkgPostgres.PGHandler, authz *pkgAuth.Authz) *IdentityRepo {
	return &IdentityRepo{
		authz: authz,
	}
}

func (_this *IdentityRepo) SavePolicy(ctx context.Context, policy *model.Policy) error {
	mapped := policy.Map2Casbin()

	err := _this.authz.AddPolicy(ctx, mapped.V0, mapped.V1, mapped.V2)
	if err != nil {
		return fmt.Errorf("error saving new policy: %v\n", err.Error())
	}
	return nil
}

func (_this *IdentityRepo) SaveGoupPolicy(ctx context.Context, group *model.GroupPolicy) error {
	mapped := group.Map2Casbin()
	return _this.authz.AddGroupingPolicy(ctx, mapped.V0, mapped.V1, mapped.V2)
}

func (_this *IdentityRepo) RemovePolicy(ctx context.Context, policy *model.Policy) error {
	mapped := policy.Map2Casbin()

	err := _this.authz.RmPolicy(ctx, mapped.V0, mapped.V1, mapped.V2)
	if err != nil {
		return fmt.Errorf("error saving new policy: %v\n", err.Error())
	}
	return nil
}

func (_this *IdentityRepo) RemoveGoupPolicy(ctx context.Context, group *model.GroupPolicy) error {
	mapped := group.Map2Casbin()
	return _this.authz.RmGroupingPolicy(ctx, mapped.V0, mapped.V1, mapped.V2)
}

func (_this *IdentityRepo) Check(ctx context.Context, sub, obj, act string) (bool, error) {
	return _this.authz.Check(ctx, sub, obj, act)
}
