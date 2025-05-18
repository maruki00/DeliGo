package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type DisableProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewDisableProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *DisableProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
