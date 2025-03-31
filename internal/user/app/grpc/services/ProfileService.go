package grpc_services

import "delivery/internal/user/infrastructure/repositories"

type ProfileService struct {
	repository repositories.UserReposiroty
}

func NewProfileService(repository repositories.UserReposiroty) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}
