package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	sharedvo "github.com/maruki00/deligo/internal/shared/value_object"
)

type SaveProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewSaveProfileHandler(repo contract.IPorofileRepository) *SaveProfileHandler {
	return &SaveProfileHandler{repo: repo}
}

func (_this *SaveProfileHandler) Handle(ctx context.Context, profile *model.Profile) error {
	return _this.repo.Save(ctx, &model.Profile{
		ID:       sharedvo.ID(profile.GetID()),
		UserID:   sharedvo.ID(profile.GetUserID()),
		FullName: profile.GetFullName(),
		Avatar:   profile.GetAvatar(),
		Bio:      profile.GetBio(),
	})
}
