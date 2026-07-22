package repository

import (
	"github.com/maruki00/deligo/internal/payment/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	UpdateStatus(id string, status models.PaymentStatus, stripeChargeID *string) error
	FindByID(id string) (*models.Payment, error)
	FindByOrderID(orderID string) (*models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) UpdateStatus(id string, status models.PaymentStatus, stripeChargeID *string) error {
	updates := map[string]interface{}{"status": status}
	if stripeChargeID != nil {
		updates["stripe_charge_id"] = *stripeChargeID
	}
	return r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *paymentRepository) FindByID(id string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) FindByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "order_id = ?", orderID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
