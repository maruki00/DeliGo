package services

import (
	"context"

	"github.com/maruki00/deligo/internal/iam/infra/repository"
)

type AuthzService struct {
	repo *repository.AuthzRepository
}

func NewAuthzService(repo *repository.AuthzRepository) *AuthzService {
	return &AuthzService{repo: repo}
}

func (_this *AuthzService) SavePolicy(ctx context.Context)
