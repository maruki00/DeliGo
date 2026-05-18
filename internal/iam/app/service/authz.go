package services

import (
	"context"
)

type AuthzService struct {
	repo *repository.AuthzRepository
}

func NewAuthzService(repo *repository.AuthzRepository) *AuthzService {
	return &AuthzService{repo: repo}
}

func (_this *AuthzService) SavePolicy(ctx context.Context) {

}
