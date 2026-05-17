package handle

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/app/command"
	"github.com/maruki00/deligo/internal/iam/infra/data/postgres"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type (
	CreatePolicy struct {
		repo *postgres.AuthzRepo
	}
	CreateGroupPolicy struct {
		repo *postgres.AuthzRepo
	}
)

func NewCreatePolicy(repo *postgres.AuthzRepo) *CreatePolicy {
	return &CreatePolicy{
		repo: repo,
	}
}

func (_this *CreatePolicy) handle(ctx context.Context, cmd pkgCqrs.Command) error {
	p := cmd.(*command.CreatePolicy).MapToPolicy()
	return _this.repo.SavePolicy(ctx, p)
}

func NewCreateGroupPolicy(repo *postgres.AuthzRepo) *CreatePolicy {
	return &CreatePolicy{
		repo: repo,
	}
}

func (_this *CreateGroupPolicy) handle(ctx context.Context, cmd pkgCqrs.Command) error {
	gp := cmd.(*command.CreateGroupPolicy).MapToGroupPolicy()
	return _this.repo.SaveGoupPolicy(ctx, gp)
}
