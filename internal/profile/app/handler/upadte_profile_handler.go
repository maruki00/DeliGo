package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
)

type UpdateProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewUpdateProfileHandler(repo contract.IPorofileRepository) *UpdateProfileHandler {
	return &UpdateProfileHandler{repo: repo}
}

func (_this *UpdateProfileHandler) Handle(ctx context.Context, id string, fields map[string]any) error {
	return _this.repo.Update(ctx, id, fields)
}
