package contracts

import (
	"context"

	aggrigate "github.com/maruki00/deligo/internal/order/domain/aggrigate"
	"github.com/maruki00/deligo/internal/order/domain/entity"
)

type IOrderRepository interface {
	Make(ctx context.Context, order *aggrigate.OrderAggrigate) (*aggrigate.OrderAggrigate, error)
	Cancel(ctx context.Context, id int) error
	Confirm(ctx context.Context, id int) error
	GetStatus(ctx context.Context, id int) (int, error)
	GetCustomerOrders(ctx context.Context, id int) ([]entity.OrderEntity, error)
	GetByFingerPrint(ctx context.Context, fingerprint string) ([]entity.OrderEntity, error)
}
