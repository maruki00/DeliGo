package service

import (
	"context"
	"fmt"

	menuDomain "github.com/maruki00/deligo/internal/shop/domain/menu"
	menuModel "github.com/maruki00/deligo/internal/shop/infra/model"
)

type MenuSVC struct {
	repo menuDomain.Repository
}

func NewMenuService(repo menuDomain.Repository) *MenuSVC {
	return &MenuSVC{repo: repo}
}

func (s *MenuSVC) CreateMenu(ctx context.Context, shopID, name string) (*menuDomain.Menu, error) {
	entity, err := menuDomain.NewMenu(fmt.Sprintf("menu-%s", shopID), shopID, name)
	if err != nil {
		return nil, err
	}

	model := &menuModel.Menu{ID: entity.ID, ShopID: entity.ShopID, Name: entity.Name, Items: []menuModel.Item{}}
	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *MenuSVC) GetMenu(ctx context.Context, id string) (*menuDomain.Menu, error) {
	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	items := make([]menuDomain.Item, 0, len(model.Items))
	for _, item := range model.Items {
		items = append(items, menuDomain.Item{ProductID: item.ProductID, Price: item.Price})
	}

	return &menuDomain.Menu{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Items: items}, nil
}

func (s *MenuSVC) ListMenusByShop(ctx context.Context, shopID string) ([]*menuDomain.Menu, error) {
	models, err := s.repo.ListByShop(ctx, shopID)
	if err != nil {
		return nil, err
	}

	items := make([]*menuDomain.Menu, 0, len(models))
	for _, model := range models {
		items = append(items, &menuDomain.Menu{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Items: []menuDomain.Item{}})
	}
	return items, nil
}

func (s *MenuSVC) AddItem(ctx context.Context, menuID, productID string, price float64) (*menuDomain.Menu, error) {
	model, err := s.repo.GetByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	entity := &menuDomain.Menu{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Items: []menuDomain.Item{}}
	for _, item := range model.Items {
		entity.Items = append(entity.Items, menuDomain.Item{ProductID: item.ProductID, Price: item.Price})
	}
	if err := entity.AddItem(productID, price); err != nil {
		return nil, err
	}

	updatedItems := make([]menuModel.Item, 0, len(entity.Items))
	for _, item := range entity.Items {
		updatedItems = append(updatedItems, menuModel.Item{ProductID: item.ProductID, Price: item.Price})
	}
	model.Items = updatedItems
	if err := s.repo.Update(ctx, model); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *MenuSVC) RemoveItem(ctx context.Context, menuID, productID string) (*menuDomain.Menu, error) {
	model, err := s.repo.GetByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	entity := &menuDomain.Menu{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Items: []menuDomain.Item{}}
	for _, item := range model.Items {
		entity.Items = append(entity.Items, menuDomain.Item{ProductID: item.ProductID, Price: item.Price})
	}
	if err := entity.RemoveItem(productID); err != nil {
		return nil, err
	}

	updatedItems := make([]menuModel.Item, 0, len(entity.Items))
	for _, item := range entity.Items {
		updatedItems = append(updatedItems, menuModel.Item{ProductID: item.ProductID, Price: item.Price})
	}
	model.Items = updatedItems
	if err := s.repo.Update(ctx, model); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *MenuSVC) DeleteMenu(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
