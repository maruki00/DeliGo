package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
)

type UpdateProfileAvatarHandler struct {
	repo contract.IPorofileRepository
}

func NewUpdateProfileAvatarHandler(repo contract.IPorofileRepository) *UpdateProfileAvatarHandler {
	return &UpdateProfileAvatarHandler{repo: repo}
}

func (_this *UpdateProfileAvatarHandler) Handle(ctx context.Context, id string, avatar string) error {
	return _this.repo.UpdateAvatar(ctx, id, avatar)
}
