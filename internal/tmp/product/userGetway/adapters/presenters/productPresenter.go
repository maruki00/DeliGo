package presenters

import (
	shared_contracts "deligo/internal/shared/domain/contracts"
	shared_models "deligo/internal/shared/infra/models"
	shared_viewmodels "deligo/internal/shared/userGateway/adapters/viewModels"
)

type ProductPresenter struct {
}

func (obj *ProductPresenter) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}

func (obj *ProductPresenter) Error(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_viewmodels.NewJsonViewModel(data)
}
