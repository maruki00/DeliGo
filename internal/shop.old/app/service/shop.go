package service

import (
	"context"
	"fmt"

	shopDomain "github.com/maruki00/deligo/internal/shop/domain/shop"
	shopModel "github.com/maruki00/deligo/internal/shop/infra/model"
)

type ShopSVC struct {
	repo shopDomain.Repository
}

func NewShopService(repo shopDomain.Repository) *ShopSVC {
	return &ShopSVC{repo: repo}
}

func (s *ShopSVC) CreateShop(ctx context.Context, name, address, phone, ownerID string) (*shopDomain.Shop, error) {
	entity, err := shopDomain.NewShop(fmt.Sprintf("shop-%d", len(name)), name, address, phone, ownerID)
	if err != nil {
		return nil, err
	}

	model := &shopModel.Shop{
		ID:      entity.ID,
		Name:    entity.Name,
		Address: entity.Address,
		Phone:   entity.Phone,
		OwnerID: entity.OwnerID,
	}

	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *ShopSVC) GetShop(ctx context.Context, id string) (*shopDomain.Shop, error) {
	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &shopDomain.Shop{ID: model.ID, Name: model.Name, Address: model.Address, Phone: model.Phone, OwnerID: model.OwnerID}, nil
}

func (s *ShopSVC) ListShops(ctx context.Context, limit, offset int) ([]*shopDomain.Shop, error) {
	models, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	items := make([]*shopDomain.Shop, 0, len(models))
	for _, model := range models {
		items = append(items, &shopDomain.Shop{ID: model.ID, Name: model.Name, Address: model.Address, Phone: model.Phone, OwnerID: model.OwnerID})
	}
	return items, nil
}

func (s *ShopSVC) UpdateShop(ctx context.Context, id, name, address, phone string) (*shopDomain.Shop, error) {
	current, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := current.UpdateDetails(name, address, phone); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, current); err != nil {
		return nil, err
	}

	return &shopDomain.Shop{ID: current.ID, Name: current.Name, Address: current.Address, Phone: current.Phone, OwnerID: current.OwnerID}, nil
}

func (s *ShopSVC) DeleteShop(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
