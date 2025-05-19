package handlers

import (
	"context"
	"deligo/internal/iam/infra/models"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type SaveProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewSaveProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *SaveProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	err := _this.repo.Save(ctx, &models.Profile{})
	return nil
}
