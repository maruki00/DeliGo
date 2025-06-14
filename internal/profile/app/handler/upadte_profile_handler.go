package handlers

import (
	"context"
	"github.com/maruki00/deligo/internal/profile/app/profile/commands"
	"github.com/maruki00/deligo/internal/profile/domain/contracts"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewUpdateProfileHandler(repo contracts.IPorofileRepository) *UpdateProfileHandler {
	return &UpdateProfileHandler{
		repo: repo,
	}
}

func (_this *UpdateProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.UpdateProfileCommand)
	return _this.repo.Update(ctx, string(cmd.ID), cmd.Fields)

}
