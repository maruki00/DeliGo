package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type Disable2ProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewDisable2ProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *Disable2ProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
