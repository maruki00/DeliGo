package services

import "github.com/maruki00/deligo/internal/iam/infra/repository"

type AuthzService struct {
	repo repository.PolicyRepository
}

func NewAuthzService(repo repository.PolicyRepository) *AuthzService {
	return &AuthzService{repo: repo}
}
