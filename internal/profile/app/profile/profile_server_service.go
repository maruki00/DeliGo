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

touch save_profile_command.go
touch delete_profile_command.go
touch upadte_profile_command.go
touch updateAvatar_profile_command.go

get_one_profile_query.go




func (_this *ProfileServerService) Save(ctx context.Context, in *profile_grpc.CreateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
func (_this *ProfileServerService) Disable(ctx context.Context, in *profile_grpc.DisableProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
func (_this *ProfileServerService) Update(ctx context.Context, in *profile_grpc.UpdateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
func (_this *ProfileServerService) UpdateAvatar(ctx context.Context, in *profile_grpc.UpdateProfileAvatareRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
func (_this *ProfileServerService) GetOne(ctx context.Context, in *profile_grpc.GETRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}