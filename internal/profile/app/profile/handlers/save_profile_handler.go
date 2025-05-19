package handlers

import (
	"context"
	"deligo/internal/profile/app/profile/commands"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/infra/models"
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	pkgCqrs "deligo/pkg/cqrs"
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
	return _this.repo.Save(ctx, &models.Profile{
		ID:       shared_valueobject.NewID(),
		UserID:   shared_valueobject.ID(cmd.UserID),
		FullName: cmd.FullName,
		Avatar:   cmd.Avatar,
	})
}
