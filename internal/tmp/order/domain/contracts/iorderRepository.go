package contracts

import (
	"context"
	aggrigate "delivery/internal/order/domain/aggrigates"
	"delivery/internal/order/domain/entities"
)

type IOrderRepository interface {
	Make(ctx context.Context, order *aggrigate.OrderAggrigate) (*aggrigate.OrderAggrigate, error)
	Cancel(ctx context.Context, id int) error
	Confirm(ctx context.Context, id int) error
	GetStatus(ctx context.Context, id int) (int, error)
	GetCustomerOrders(ctx context.Context, id int) ([]entities.OrderEntity, error)
	GetByFingerPrint(ctx context.Context, fingerprint string) ([]entities.OrderEntity, error)
}
