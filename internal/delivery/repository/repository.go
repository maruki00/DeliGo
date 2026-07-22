package repository

import (
	"time"

	"github.com/maruki00/deligo/delivery/models"

	"gorm.io/gorm"
)

type DeliveryRepository interface {
	GetActiveCouriers() ([]models.Courier, error)
	GetCourierBusyStatus(courierID string) (bool, error)
	CreateOrderCourier(assignment *models.OrderCourier) error
	UpdateOrderCourier(assignment *models.OrderCourier) error
	GetOrderCourierByOrder(orderID string) (*models.OrderCourier, error)
	UpdateCourierLocation(courierID string, lat, lon float64) error
	SaveTrackingLog(log *models.OrderTracking) error
	CreateCourier(courier *models.Courier) error // For setup/testing purposes
}

type deliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) DeliveryRepository {
	return &deliveryRepository{db: db}
}

func (r *deliveryRepository) GetActiveCouriers() ([]models.Courier, error) {
	var couriers []models.Courier
	err := r.db.Where("is_active = ?", true).Find(&couriers).Error
	return couriers, err
}

func (r *deliveryRepository) GetCourierBusyStatus(courierID string) (bool, error) {
	var count int64
	activeStatuses := []string{"accepted", "at_restaurant", "picked_up"}
	err := r.db.Model(&models.OrderCourier{}).
		Where("courier_id = ? AND status IN ?", courierID, activeStatuses).
		Count(&count).Error
	return count > 0, err
}

func (r *deliveryRepository) CreateOrderCourier(assignment *models.OrderCourier) error {
	return r.db.Create(assignment).Error
}

func (r *deliveryRepository) UpdateOrderCourier(assignment *models.OrderCourier) error {
	return r.db.Save(assignment).Error
}

func (r *deliveryRepository) GetOrderCourierByOrder(orderID string) (*models.OrderCourier, error) {
	var assignment models.OrderCourier
	err := r.db.Where("order_id = ?", orderID).First(&assignment).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *deliveryRepository) UpdateCourierLocation(courierID string, lat, lon float64) error {
	return r.db.Model(&models.Courier{}).
		Where("id = ?", courierID).
		Updates(map[string]interface{}{
			"current_latitude":  lat,
			"current_longitude": lon,
			"updated_at":        time.Now(),
		}).Error
}

func (r *deliveryRepository) SaveTrackingLog(log *models.OrderTracking) error {
	return r.db.Create(log).Error
}

func (r *deliveryRepository) CreateCourier(courier *models.Courier) error {
	return r.db.Create(courier).Error
}
