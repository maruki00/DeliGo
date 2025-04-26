package ports

import (
	shared_contracts "deligo/internal/shared/domain/contracts"
	shared_models "deligo/internal/shared/infra/models"
)

type ProductOutputPort interface {
	Success(data shared_models.ResponseModel) shared_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_contracts.ViewModel
}
