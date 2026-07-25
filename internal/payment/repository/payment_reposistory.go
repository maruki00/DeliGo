package repository

import (
	models "github.com/maruki00/deligo/internal/payment/domain"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) UpdateStatus(id string, status models.PaymentStatus, stripeChargeID *string) error {
	updates := map[string]interface{}{"status": status}
	if stripeChargeID != nil {
		updates["stripe_charge_id"] = *stripeChargeID
	}
	return r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PaymentRepository) FindByID(id string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) FindByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "order_id = ?", orderID).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
