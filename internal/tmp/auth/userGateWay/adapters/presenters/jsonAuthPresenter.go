package presenters

import (
	"deligo/internal/auth/userGateWay/adapters/viewmodels"
	shared_contracts "deligo/internal/shared/domain/contracts"
	shared_models "deligo/internal/shared/infra/models"
)

type JsonAuthPresenter struct {
}

func NewJSONAuthPresenter() JsonAuthPresenter {
	return JsonAuthPresenter{}
}

func (obj *JsonAuthPresenter) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return viewmodels.NewJsonViewModel(data)
}

func (obj *JsonAuthPresenter) Error(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return viewmodels.NewJsonViewModel(data)
}
