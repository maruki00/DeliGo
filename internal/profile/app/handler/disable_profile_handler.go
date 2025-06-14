package handlers

import (
	"context"
	"github.com/maruki00/deligo/internal/profile/app/profile/commands"
	"github.com/maruki00/deligo/internal/profile/domain/contracts"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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
	return _this.repo.Disable(ctx, string(command.(commands.DiscableProfileCommand).ID))
}
