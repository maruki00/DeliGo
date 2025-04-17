package grpc_services

import (
	"context"
	profile_grpc "delivery/internal/user/infra/grpc/profile"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type ProfileService struct {
	repository *repositories.ProfileRepository
	profile_grpc.UnimplementedProfileServiceServer
}

func NewProfileService(repository *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}

func (_this *ProfileService) Create(ctx context.Context, in *profile_grpc.CreateProfileRequest) (*profile_grpc.ProfileResponse, error) {
	res, err := _this.repository.Create(ctx, &models.Profile{
		UserID:   in.UserID,
		FullName: in.FullName,
		Avatar:   in.Avatar,
		Bio:      in.Bio,
	})
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":        res.GetID(),
		"Full_name": res.GetFullName(),
		"bio":       res.GetBio(),
		"avatr":     res.GetAvatar(),
	})

	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  []*structpb.Value{stuctRes},
	}, nil

}
func (_this *ProfileService) Delete(context.Context, *profile_grpc.DeleteProfileRequest) (*profile_grpc.ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (_this *ProfileService) Update(context.Context, *profile_grpc.UpdateProfileRequest) (*profile_grpc.ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (_this *ProfileService) GetOne(context.Context, *profile_grpc.EmptyProfileResponse) (*profile_grpc.ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (_this *ProfileService) GetMany(context.Context, *profile_grpc.EmptyProfileResponse) (*profile_grpc.ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (_this *ProfileService) Search(context.Context, *profile_grpc.EmptyProfileResponse) (*profile_grpc.ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
