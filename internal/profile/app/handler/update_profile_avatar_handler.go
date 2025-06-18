package handler

import (
	"context"
	"github.com/maruki00/deligo/internal/profile/app/profile/commands"
	"github.com/maruki00/deligo/internal/profile/domain/contracts"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateProfileAvatarHandler struct {
	repo contracts.IPorofileRepository
}

func NewUpdateProfileAvatarHandler(repo contracts.IPorofileRepository) *UpdateProfileAvatarHandler {
	return &UpdateProfileAvatarHandler{
		repo: repo,
	}
}

func (_this *UpdateProfileAvatarHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.UpdateProfileAvatarCommand)
	return _this.repo.UpdateAvatar(ctx, string(cmd.ID), cmd.Avatar)
}
