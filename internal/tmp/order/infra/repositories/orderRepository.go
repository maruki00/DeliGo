package repositories

import (
	"context"
	aggrigate "delivery/internal/order/domain/aggrigates"
	"delivery/internal/order/domain/entities"
	"delivery/internal/order/infra/models"

	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewOrderRepository(db *gorm.DB, model interface{}) *OrderRepository {
	return &OrderRepository{
		db:    db,
		model: model,
	}
}

func (obj *OrderRepository) Make(ctx context.Context, order *aggrigate.OrderAggrigate) (*aggrigate.OrderAggrigate, error) {

	obj.db.Begin()
	res := obj.db.WithContext(ctx).Model(obj.model).Create(order.Order)
	fmt.Println(order.Order)
	if res.Error != nil {
		obj.db.Rollback()
		return nil, errors.New("could not create the order")
	}

	d := obj.db.Model(&models.OrderProduct{})
	for _, product := range order.Items {
		res := d.Create(&product)
		if res.Error != nil {
			obj.db.Rollback()
			return nil, errors.New("could not create order prodcut")
		}
	}

	obj.db.Commit()
	return order, nil
}

func (obj *OrderRepository) Cancel(ctx context.Context, id int) error {
	res := obj.db.WithContext(ctx).Where("id = ?", id).Delete(obj.model).Update("is_cancelled", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}

func (obj *OrderRepository) Confirm(ctx context.Context, id int) error {
	res := obj.db.WithContext(ctx).Where("id = ?", id).Update("is_confirmed", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}

func (obj *OrderRepository) GetStatus(ctx context.Context, id int) (int, error) {
	var order models.Order
	res := obj.db.WithContext(ctx).Model(order).Where("id = ?", id).Find(&order)
	if res.Error != nil {
		return -1, res.Error
	}
	return order.Status, nil
}

func (obj *OrderRepository) GetCustomerOrders(ctx context.Context, id int) ([]entities.OrderEntity, error) {
	var orders []entities.OrderEntity
	res := obj.db.WithContext(ctx).Model(obj.model).Where("customer_id = ?", id).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}

func (obj *OrderRepository) GetByFingerPrint(ctx context.Context, fingerprint string) ([]entities.OrderEntity, error) {
	orders := []entities.OrderEntity{}
	res := obj.db.WithContext(ctx).Model(obj.model).Where("fingerprint = ?", fingerprint).Find(orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}
