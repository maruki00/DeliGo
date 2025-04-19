package ports

import (
	shared_domain_contracts "deligo/internal/shared/domain/contracts"
	shared_models "deligo/internal/shared/infra/models"
)

type AuthOutputPort interface {
	Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
}
