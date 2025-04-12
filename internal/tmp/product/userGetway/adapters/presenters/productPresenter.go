package presenters

import (
	shared_contracts "delivery/internal/shared/domain/contracts"
	shared_models "delivery/internal/shared/infra/models"
	shared_viewmodels "delivery/internal/shared/userGateway/adapters/viewModels"
)

type ProductPresenter struct {
}

func (obj *ProductPresenter) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}

func (obj *ProductPresenter) Error(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}
