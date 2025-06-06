package userHandler

import (
	"context"
	userQueries "deligo/internal/iam/app/user/queries"
	"deligo/internal/iam/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
	"fmt"
)

type FindUserByUsernameHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByUsernameHandler(userRepo contracts.IUserRepository) *FindUserByUsernameHandler {
	return &FindUserByUsernameHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByUsernameHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	fmt.Println("called : FindUserByUsernameHandler")
	cmd := query.(*userQueries.FindUserByUsernameQuery)
	user, err := _this.userRepo.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
