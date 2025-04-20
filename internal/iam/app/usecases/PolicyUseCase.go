package usecases

import "deligo/internal/iam/domain/contracts"

type PolicyUseCase struct {
	policyRepo contracts.IPolicyRepository
}

func NewPolicyUseCase(policyRepo contracts.IPolicyRepository) *PolicyUseCase {
	return &PolicyUseCase{
		policyRepo: policyRepo,
	}
}
