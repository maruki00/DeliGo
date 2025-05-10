package grpc_services

import (
	"context"
	"deligo/internal/profile/domain/contracts"
	profile_grpc "deligo/internal/profile/infra/grpc/profile"

	"google.golang.org/grpc"
)

type ProfileServerService struct {
	repository contracts.IPorofileRepository
	profile_grpc.UnimplementedProfileServiceServer
}

func NewProfileService(repository contracts.IPorofileRepository) *ProfileServerService {
	return &ProfileServerService{
		repository: repository,
	}
}

Save
Delete
Upadte
UpdateAvatar
func (_this *ProfileServerService) Save(ctx context.Context, in *profile_grpc.CreateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}

func (_this *ProfileServerService) Delete(ctx context.Context, in *profile_grpc.DeleteProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}

func (_this *ProfileServerService) Update(ctx context.Context, in *profile_grpc.UpdateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}

func (_this *ProfileServerService) UpdateAvatar(ctx context.Context, in *profile_grpc.UpdateProfileAvatareRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}

func (_this *ProfileServerService) UpdatePassword(ctx context.Context, in *profile_grpc.UpdateProfilePasswordRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}

func (_this *ProfileServerService) GetOne(ctx context.Context, in *profile_grpc.GETRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
