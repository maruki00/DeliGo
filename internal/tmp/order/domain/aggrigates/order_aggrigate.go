package aggrigate

import "delivery/internal/order/domain/entities"

type OrderAggrigate struct {
	Order entities.OrderEntity
	Items []interface{}
	Price float32
}
