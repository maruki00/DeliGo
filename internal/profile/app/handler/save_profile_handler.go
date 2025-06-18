package handler

import (
	"context"

	command "github.com/maruki00/deligo/internal/profile/app/commnd"
	"github.com/maruki00/deligo/internal/profile/domain/contract"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type SaveProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewSaveProfileHandler(repo contract.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *SaveProfileHandler) Handle(ctx context.Context, c pkgCqrs.Command) error {
	cmd := c.(*command.SaveProfileCommand)
	return _this.repo.Save(ctx, &model.Profile{
		ID:       sharedvo.NewID(),
		UserID:   sharedvo.ID(cmd.UserID),
		FullName: cmd.FullName,
		Avatar:   cmd.Avatar,
	})
}
