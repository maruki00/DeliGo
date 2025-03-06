package services

import (
	"delivery/internal/shop/domain/contracts"
	shop_domain_contracts "delivery/internal/shop/domain/contracts"
)

type ShopService struct {
	repo shop_domain_contracts.IShopRepository
}

func NewShopService(repo contracts.IShopRepository) *ShopService {
	return &ShopService{
		repo: repo,
	}
}

func (obj *ShopService) OpenShop() error {

	return nil
}

func (obj *ShopService) CloseShop() error {

	return nil
}

func (obj *ShopService) UpdateShopInfo() error {
	return nil
}

func (obj *ShopService) GetShopStatus() (string, error) {
	status := "open"
	return status, nil
}
