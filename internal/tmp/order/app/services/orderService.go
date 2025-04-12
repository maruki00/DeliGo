package services

import (
	aggrigate "delivery/internal/order/domain/aggrigates"
	"delivery/internal/order/domain/contracts"
	order_domain_contracts "delivery/internal/order/domain/contracts"
	"delivery/internal/order/domain/dtos"
	"delivery/internal/order/domain/ports"
	order_domain_ports "delivery/internal/order/domain/ports"
	shared_contracts "delivery/internal/shared/domain/contracts"
)

type OrderService struct {
	repository contracts.IOrderRepository
	outputPort ports.OrderOutputPort
}

func NewOrderService(
	repository order_domain_contracts.IOrderRepository,
	outputPort order_domain_ports.OrderOutputPort,
) *OrderService {
	return &OrderService{
		repository: repository,
		outputPort: outputPort,
	}
}

func (obj *OrderService) CreateOrder(dto dtos.CreateNewOrderDTO) shared_contracts.ViewModel {

	// orderAggrigate, err := order_factories.NewOrderAggrigate(dto.CostomerId, dto.Products)
	// if err != nil {
	// 	panic(err)
	// }

	// obj.repository.Make(orderAggrigate)
	return nil
}

func (obj *OrderService) Make(order *aggrigate.OrderAggrigate) shared_contracts.ViewModel {

	return nil
}

func (obj *OrderService) Cancel(id int) shared_contracts.ViewModel {
	return nil
}

func (obj *OrderService) Confirm(id int) shared_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetStatus(id int) shared_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetCustomerOrders(id int) shared_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetByFingerPrint(fingerprint string) shared_contracts.ViewModel {
	return nil
}
