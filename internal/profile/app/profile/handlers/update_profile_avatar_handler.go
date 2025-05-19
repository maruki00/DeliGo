package handlers

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	pkgCqrs "deligo/pkg/cqrs"
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

	return nil
}
