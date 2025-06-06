package aggrigate

import "deligo/internal/order/domain/entity"

type OrderAggrigate struct {
	Order entity.OrderEntity
	Items []entity.OrderProductEntity
	Price float32
}
