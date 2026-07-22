package repository

import (
	"context"
	"errors"

	"github.com/maruki00/deligo/order/domain"

	"gorm.io/gorm"
)

type postgresOrderRepository struct {
	db *gorm.DB
}

func NewPostgresOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &postgresOrderRepository{db: db}
}

func (r *postgresOrderRepository) Create(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(order).Error
	})
}

func (r *postgresOrderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.WithContext(ctx).Preload("Items").First(&order, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("order not found")
	}
	return &order, err
}

func (r *postgresOrderRepository) UpdateStatus(ctx context.Context, id string, currentStatus, newStatus domain.OrderStatus) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&domain.Order{}).Where("id = ?", id)
		if currentStatus != "" {
			query = query.Where("status = ?", currentStatus)
		}

		result := query.Update("status", newStatus)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("order status transition invalid or order not found")
		}
		return nil
	})
}

func (r *postgresOrderRepository) Update(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}
