package ports

import (
	shared_domain_contracts "delivery/internal/shared/domain/contracts"
	shared_models "delivery/internal/shared/infra/models"
)

type AuthOutputPort interface {
	Success(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
	Error(data shared_models.ResponseModel) shared_domain_contracts.ViewModel
}
