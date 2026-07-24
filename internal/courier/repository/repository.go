package repository

import (
	"time"

	"github.com/maruki00/deligo/internal/courier/domain"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (r *DeliveryRepository) GetActiveCouriers() ([]domain.Courier, error) {
	var couriers []domain.Courier
	err := r.db.Where("is_active = ?", true).Find(&couriers).Error
	return couriers, err
}

func (r *DeliveryRepository) GetCourierBusyStatus(courierID string) (bool, error) {
	var count int64
	activeStatuses := []string{"accepted", "at_restaurant", "picked_up"}
	err := r.db.Model(&domain.OrderCourier{}).
		Where("courier_id = ? AND status IN ?", courierID, activeStatuses).
		Count(&count).Error
	return count > 0, err
}

func (r *DeliveryRepository) CreateOrderCourier(assignment *domain.OrderCourier) error {
	return r.db.Create(assignment).Error
}

func (r *DeliveryRepository) UpdateOrderCourier(assignment *domain.OrderCourier) error {
	return r.db.Save(assignment).Error
}

func (r *DeliveryRepository) GetOrderCourierByOrder(orderID string) (*domain.OrderCourier, error) {
	var assignment domain.OrderCourier
	err := r.db.Where("order_id = ?", orderID).First(&assignment).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *DeliveryRepository) UpdateCourierLocation(courierID string, lat, lon float64) error {
	return r.db.Model(&domain.Courier{}).
		Where("id = ?", courierID).
		Updates(map[string]interface{}{
			"current_latitude":  lat,
			"current_longitude": lon,
			"updated_at":        time.Now(),
		}).Error
}

func (r *DeliveryRepository) SaveTrackingLog(log *domain.OrderTracking) error {
	return r.db.Create(log).Error
}

func (r *DeliveryRepository) CreateCourier(courier *domain.Courier) error {
	return r.db.Create(courier).Error
}
