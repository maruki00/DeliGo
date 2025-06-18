package service

import (
	contracts "github.com/maruki00/deligo/internal/order/domain/contract"
)

type OrderService struct {
	repo contracts.IOrderRepository
}

func NewOrderService(
	repo contracts.IOrderRepository,
) *OrderService {
	return &OrderService{
		repo: repo,
	}
}
