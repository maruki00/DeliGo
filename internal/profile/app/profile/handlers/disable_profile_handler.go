package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/infra/models"
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
	_this.repo.Disable(ctx, &models.Profile{
	ID        : command.,
	UserID    : command.,
	FullName  : command.,
	Avatar    : command.,
	Bio       : command.,
	})
	return nil
}
