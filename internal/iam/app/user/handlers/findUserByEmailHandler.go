package userHandlers

import (
	"context"
	userQueries "deligo/internal/iam/app/user/queries"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
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
