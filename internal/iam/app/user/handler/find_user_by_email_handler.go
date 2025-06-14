package userHandler

import (
	"context"
	userQueries "github.com/maruki00/deligo/internal/iam/app/user/queries"
	"github.com/maruki00/deligo/internal/iam/domain/contracts"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type FindUserByEmailHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByEmailHandler(userRepo contracts.IUserRepository) *FindUserByEmailHandler {
	return &FindUserByEmailHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByEmailHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	qry := query.(*userQueries.FindUserByEmailQuery)
	user, err := _this.userRepo.FindByEmail(ctx, qry.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
