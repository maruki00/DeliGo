package service

import (
	"context"
	"fmt"

	productDomain "github.com/maruki00/deligo/internal/shop/domain/product"
)

type ProductSVC struct {
	repo productDomain.Repository
}

func NewProductService(repo productDomain.Repository) *ProductSVC {
	return &ProductSVC{repo: repo}
}

func (s *ProductSVC) CreateProduct(ctx context.Context, shopID, name, description string, price float64, stock int) (*productDomain.Product, error) {
	entity, err := productDomain.NewProduct(fmt.Sprintf("product-%s", shopID), shopID, name, description, price, stock)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *ProductSVC) GetProduct(ctx context.Context, id string) (*productDomain.Product, error) {
	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &productDomain.Product{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Description: model.Description, Price: model.Price, Stock: model.Stock}, nil
}

func (s *ProductSVC) ListProductsByShop(ctx context.Context, shopID string, limit, offset int) ([]*productDomain.Product, error) {
	models, err := s.repo.ListByShop(ctx, shopID, limit, offset)
	if err != nil {
		return nil, err
	}

	items := make([]*productDomain.Product, 0, len(models))
	for _, model := range models {
		items = append(items, &productDomain.Product{ID: model.ID, ShopID: model.ShopID, Name: model.Name, Description: model.Description, Price: model.Price, Stock: model.Stock})
	}
	return items, nil
}

func (s *ProductSVC) UpdateProduct(ctx context.Context, id, name, description string, price float64) (*productDomain.Product, error) {
	current, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := current.UpdateDetails(name, description, price); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, current); err != nil {
		return nil, err
	}

	return &productDomain.Product{ID: current.ID, ShopID: current.ShopID, Name: current.Name, Description: current.Description, Price: current.Price, Stock: current.Stock}, nil
}

func (s *ProductSVC) AdjustStock(ctx context.Context, id string, delta int) (*productDomain.Product, error) {
	current, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := current.AdjustStock(delta); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, current); err != nil {
		return nil, err
	}

	return &productDomain.Product{ID: current.ID, ShopID: current.ShopID, Name: current.Name, Description: current.Description, Price: current.Price, Stock: current.Stock}, nil
}

func (s *ProductSVC) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
