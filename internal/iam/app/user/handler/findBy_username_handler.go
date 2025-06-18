package userHandler

import (
	"context"
	"fmt"

	userQueries "github.com/maruki00/deligo/internal/iam/app/user/query"
	"github.com/maruki00/deligo/internal/iam/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type FindUserByUsernameHandler struct {
	userRepo contract.IUserRepository
}

func NewFindUserByUsernameHandler(userRepo contract.IUserRepository) *FindUserByUsernameHandler {
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
