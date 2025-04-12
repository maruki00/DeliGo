package factories

import "delivery/internal/order/app/services"

func NewOrderService() *services.OrderService {
	return &services.OrderService{}
}
