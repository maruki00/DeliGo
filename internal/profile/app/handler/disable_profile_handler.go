package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
)

type DisableProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewDisableProfileHandler(repo contract.IPorofileRepository) *DisableProfileHandler {
	return &DisableProfileHandler{repo: repo}
}

func (_this *DisableProfileHandler) Handle(ctx context.Context, id string) error {
	return _this.repo.Disable(ctx, id)
}
