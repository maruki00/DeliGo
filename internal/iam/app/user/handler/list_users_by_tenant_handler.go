package userHandler

import (
	"context"

	userQueries "github.com/maruki00/deligo/internal/iam/app/user/query"
	"github.com/maruki00/deligo/internal/iam/domain/contract"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type ListUsersByTenantHandler struct {
	userRepo contract.IUserRepository
}

func NewListUsersByTenantHandler(userRepo contract.IUserRepository) *ListUsersByTenantHandler {
	return &ListUsersByTenantHandler{
		userRepo: userRepo,
	}
}

func (_this *ListUsersByTenantHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	qry := query.(*userQueries.ListUsersByTenantQuery)
	users, err := _this.userRepo.ListByTenant(ctx, sharedvo.ID(qry.TenantID.String()), qry.Pagination)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, nil
	}
	return users, nil
}
