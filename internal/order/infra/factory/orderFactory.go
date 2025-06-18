package factory

import "github.com/maruki00/deligo/internal/order/app/service"


func NewOrderService() *service.OrderService {
	return &service.OrderService{}
}
