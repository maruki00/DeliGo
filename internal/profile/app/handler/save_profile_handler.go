package handlers

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/app/profile/commands"
	"github.com/maruki00/deligo/internal/profile/domain/contracts"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	shared_valueobject "github.com/maruki00/deligo/internal/shared/domain/sharedvo"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type SaveProfileHandler struct {
	repo contracts.IPorofileRepository
}

func NewSaveProfileHandler(repo contracts.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{
		repo: repo,
	}
}

func (_this *SaveProfileHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.SaveProfileCommand)
	return _this.repo.Save(ctx, &model.Profile{
		ID:       shared_valueobject.NewID(),
		UserID:   shared_valueobject.ID(cmd.UserID),
		FullName: cmd.FullName,
		Avatar:   cmd.Avatar,
	})
}
