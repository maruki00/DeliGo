package userHandler

import (
	"context"
	userQueries "github.com/maruki00/deligo/internal/iam/app/user/queries"
	"github.com/maruki00/deligo/internal/iam/domain/contracts"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type FindUserByIdHandler struct {
	userRepo contracts.IUserRepository
}

func NewFindUserByIdHandler(userRepo contracts.IUserRepository) *FindUserByIdHandler {
	return &FindUserByIdHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByIdHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	qry := query.(*userQueries.FindUserByIdQuery)
	user, err := _this.userRepo.FindByID(ctx, valueobjects.ID(qry.ID.String()))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
