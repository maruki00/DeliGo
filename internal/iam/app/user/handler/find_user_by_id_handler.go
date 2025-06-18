package userHandler

import (
	"context"

	userQueries "github.com/maruki00/deligo/internal/iam/app/user/query"
	"github.com/maruki00/deligo/internal/iam/domain/contract"
	valueobjects "github.com/maruki00/deligo/internal/iam/domain/valueobject"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type FindUserByIdHandler struct {
	userRepo contract.IUserRepository
}

func NewFindUserByIdHandler(userRepo contract.IUserRepository) *FindUserByIdHandler {
	return &FindUserByIdHandler{
		userRepo: userRepo,
	}
}

func (_this *FindUserByIdHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {
	qry := query.(*userQueries.FindUserByIdQuery)
	user, err := _this.userRepo.FindByID(ctx, sharedvo.ID(qry.ID.String()))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
