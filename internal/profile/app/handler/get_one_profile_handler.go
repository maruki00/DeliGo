package handler

import (
	"context"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
	"github.com/maruki00/deligo/internal/profile/infra/model"
)

type GetOneProfileHandler struct {
	repo contract.IPorofileRepository
}

func NewGetOneProfileHandler(repo contract.IPorofileRepository) *GetOneProfileHandler {
	return &GetOneProfileHandler{repo: repo}
}

func (_this *GetOneProfileHandler) Handle(ctx context.Context, id string) (*model.Profile, error) {
	return _this.repo.FindByID(ctx, id)
}
