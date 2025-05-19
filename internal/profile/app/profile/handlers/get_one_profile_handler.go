package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetOneProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewGetOneProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *GetOneProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
