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




Save(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
Disable(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
Update(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
UpdateAvatar(ctx context.Context, in *UpdateProfileAvatareRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
GetOne(ctx context.Context, in *GETRequest, opts ...grpc.CallOption) (*ProfileResponse, error)