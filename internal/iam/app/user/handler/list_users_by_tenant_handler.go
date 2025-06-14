package userHandler

import (
	"context"
	userQueries "github.com/maruki00/deligo/internal/iam/app/user/queries"
	"github.com/maruki00/deligo/internal/iam/domain/contracts"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type ListUsersByTenantHandler struct {
	userRepo contracts.IUserRepository
}

func NewListUsersByTenantHandler(userRepo contracts.IUserRepository) *ListUsersByTenantHandler {
	return &ListUsersByTenantHandler{
		userRepo: userRepo,
	}
}

func (_this *ListUsersByTenantHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	qry := query.(*userQueries.ListUsersByTenantQuery)
	users, err := _this.userRepo.ListByTenant(ctx, valueobjects.ID(qry.TenantID.String()), qry.Pagination)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, nil
	}
	return users, nil
}
