package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type Disable1ProfileHandler struct {
	repo contracts.IPorofileRepository
}

func Disable1ProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *Disable1ProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
