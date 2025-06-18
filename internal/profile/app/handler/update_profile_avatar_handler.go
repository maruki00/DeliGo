package handler

import (
	"context"

	command "github.com/maruki00/deligo/internal/profile/app/commnd"
	"github.com/maruki00/deligo/internal/profile/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateProfileAvatarHandler struct {
	repo contract.IPorofileRepository
}

func NewUpdateProfileAvatarHandler(repo contract.IPorofileRepository) *UpdateProfileAvatarHandler {
	return &UpdateProfileAvatarHandler{
		repo: repo,
	}
}

func (_this *UpdateProfileAvatarHandler) Handle(ctx context.Context, c pkgCqrs.Command) error {
	cmd := c.(*command.UpdateProfileAvatarCommand)
	return _this.repo.UpdateAvatar(ctx, string(cmd.ID), cmd.Avatar)
}
