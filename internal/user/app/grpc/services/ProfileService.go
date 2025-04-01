package grpc_services

import (
	"context"
	profile_grpc "delivery/internal/user/infra/grpc/profile"
	"delivery/internal/user/infra/repositories"
)

type ProfileService struct {
	repository repositories.ProfileRepository
	profile_grpc.UnimplementedProfileServiceServer
}

func NewProfileService(repository repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}

func (ps *ProfileService) Create(ctx context.Context, payload *profile_grpc.CreateProfileRequest) (*profile_grpc.ProfileResponse, error) {
	p := &profile_grpc.CreateProfileRequest{
		ID:       "12344",     //
		UserID:   "123215234", //
		FullName: "123215234", //
		Avatar:   "123215234", //
		Bio:      "123215234", //
	}
	res, err := ps.UnimplementedProfileServiceServer.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
