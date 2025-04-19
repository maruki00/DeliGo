package aggrigate

import "deligo/internal/order/domain/entities"

type OrderAggrigate struct {
	Order entities.OrderEntity
	Items []interface{}
	Price float32
}
