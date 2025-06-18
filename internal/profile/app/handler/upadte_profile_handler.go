package handler

import (
	"context"

	command "github.com/maruki00/deligo/internal/profile/app/commnd"
	"github.com/maruki00/deligo/internal/profile/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type UpdateProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewUpdateProfileHandler(repo contract.IPorofileRepository) *UpdateProfileHandler {
	return &UpdateProfileHandler{
		repo: repo,
	}
}

func (_this *UpdateProfileHandler) Handle(ctx context.Context, c pkgCqrs.Command) error {
	cmd := c.(*command.UpdateProfileAvatarCommand)
	//TODO: add fields
	return _this.repo.Update(ctx, string(cmd.ID), nil)

}
