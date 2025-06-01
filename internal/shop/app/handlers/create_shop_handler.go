package handlers

import (
	"context"
	shared_models "deligo/internal/shared/infra/models"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/domain/contracts"
	"deligo/internal/shop/infra/models"
	pkgCqrs "deligo/pkg/cqrs"
)

type CreateShopHHandler struct {
	repo contracts.IShopRepository
}

func NewCreateShopHandler(repo contracts.IShopRepository) *CreateShopHHandler {
	return &CreateShopHHandler{
		repo: repo,
	}
}

func (_this *CreateShopHHandler) handler(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*commands.CreateShopCommand)

	return _this.repo.Save(ctx, &models.Shop{
		BaseModel: shared_models.BaseModel{
			ID: cmd.ID,
		},
		Name:    cmd.ShopName,
		OpenAt:  cmd.OpenAt,
		CloseAt: cmd.CloseAt,
	})
}
