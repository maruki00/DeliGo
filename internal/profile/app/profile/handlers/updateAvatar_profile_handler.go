package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type DisableProfileHandler4 struct {
	repo contracts.IPorofileRepository
}

func NewDisableProfileHandler4(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *DisableProfileHandler4) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
