package service

import (
	aggrigate "github.com/maruki00/deligo/internal/order/domain/aggrigate"
	contracts "github.com/maruki00/deligo/internal/order/domain/contract"
	"github.com/maruki00/deligo/internal/order/domain/dtos"
	order_domain_ports "github.com/maruki00/deligo/internal/order/domain/port"
	shared_contracts "github.com/maruki00/deligo/internal/shared/domain/contract"
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
