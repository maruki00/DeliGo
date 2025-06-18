package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type GetOneProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewGetOneProfileHandler(repo contract.IPorofileRepository) *GetOneProfileHandler {
	return &GetOneProfileHandler{
		repo: repo,
	}
}
func (_this *GetOneProfileHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
