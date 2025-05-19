package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
)

type DisableProfileHandler3struct {
	repo contracts.IPorofileRepository
}

func NewDisableProfileHandler3(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *DisableProfileHandler3) Handle(ctx context.Context, command pkgCqrs.Command) error {

	return nil
}
