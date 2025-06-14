package repository

import (
	"context"
	aggrigate "github.com/maruki00/deligo/internal/order/domain/aggrigate"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *pkgPostgres.PGHandler
}

func NewOrderRepository(db *pkgPostgres.PGHandler) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

/**

Create
UpdateBillingAddress
UpdateShippingAddress
UpdateStatus
GetOrder

*/

func (_this *OrderRepository) Create(ctx context.Context, order *aggrigate.OrderAggrigate) error {
	return _this.db.GetDB().Transaction(func(tx *gorm.DB) error {

	})
}

func (_this *OrderRepository) UpdateShippingAddress(ctx context.Context) error {
	return nil
}

func (_this *OrderRepository) UpdateBillingAddress(ctx context.Context) error {

	return nil
}

func (_this *OrderRepository) UpdateStatus(ctx context.Context) error {

	return nil
}

func (_this *OrderRepository) GetOrder(ctx context.Context) error {

	return nil
}
