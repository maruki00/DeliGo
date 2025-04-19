package factories

import "deligo/internal/order/app/services"

func NewOrderService() *services.OrderService {
	return &services.OrderService{}
}
