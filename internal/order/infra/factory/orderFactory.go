package factories

import "github.com/maruki00/deligo/internal/order/app/services"

func NewOrderService() *services.OrderService {
	return &services.OrderService{}
}
