package handler

import (
	"context"

	command "github.com/maruki00/deligo/internal/profile/app/commnd"
	"github.com/maruki00/deligo/internal/profile/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type DisableProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewDisableProfileHandler(repo contract.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *DisableProfileHandler) Handle(ctx context.Context, cmd pkgCqrs.Command) error {
	return _this.repo.Disable(
		ctx,
		cmd.(*command.DiscableProfileCommand).ID.String(),
	)
}
