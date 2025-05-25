package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetOneProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewGetOneProfileHandler(repo contracts.IPorofileRepository) *GetOneProfileHandler {
	return &GetOneProfileHandler{
		repo: repo,
	}
}
func (_this *GetOneProfileHandler) Handle(ctx context.Context, query pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
