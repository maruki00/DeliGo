package service

import (
	"context"
	"net/http"

	profile_contract "github.com/maruki00/deligo/internal/profile/domain/contract"
	profile_grpc "github.com/maruki00/deligo/internal/profile/infra/grpc/profile"
	"github.com/maruki00/deligo/internal/profile/infra/model"
	shared_valueobject "github.com/maruki00/deligo/internal/shared/value_object"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type ProfileServerService struct {
	profile_grpc.UnimplementedProfileServiceServer
	repo profile_contract.IPorofileRepository
}

func NewProfileService(repo profile_contract.IPorofileRepository) *ProfileServerService {
	return &ProfileServerService{repo: repo}
}

func CommandCheck(err error) (*profile_grpc.ProfileResponse, error) {
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Result:  nil,
		}, err
	}
	return &profile_grpc.ProfileResponse{
		Code:    http.StatusAccepted,
		Message: "success",
		Result:  nil,
	}, nil
}

func (_this *ProfileServerService) Save(ctx context.Context, in *profile_grpc.CreateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	err := _this.repo.Save(ctx, &model.Profile{
		ID:       shared_valueobject.NewID(),
		UserID:   shared_valueobject.ID(in.UserID),
		FullName: in.FullName,
		Avatar:   in.Avatar,
		Bio:      in.Bio,
	})
	return CommandCheck(err)
}

func (_this *ProfileServerService) Disable(ctx context.Context, in *profile_grpc.DisableProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	err := _this.repo.Disable(ctx, string(shared_valueobject.ID(in.ID)))
	return CommandCheck(err)
}

func (_this *ProfileServerService) Update(ctx context.Context, in *profile_grpc.UpdateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	fields := make(map[string]any, len(in.Fields))
	for k, v := range in.Fields {
		fields[k] = v
	}
	err := _this.repo.Update(ctx, string(shared_valueobject.ID(in.ID)), fields)
	return CommandCheck(err)
}

func (_this *ProfileServerService) UpdateAvatar(ctx context.Context, in *profile_grpc.UpdateProfileAvatareRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	err := _this.repo.UpdateAvatar(ctx, string(shared_valueobject.ID(in.ID)), in.Avatar)
	return CommandCheck(err)
}

func (_this *ProfileServerService) GetOne(ctx context.Context, in *profile_grpc.GETRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	id := ""
	if in.QueryParams != nil {
		if value, ok := in.QueryParams.Fields["id"]; ok {
			if valueStr, ok := value.AsInterface().(string); ok {
				id = valueStr
			}
		}
	}
	if id == "" {
		return &profile_grpc.ProfileResponse{
			Code:    http.StatusBadRequest,
			Message: "id is required",
			Result:  nil,
		}, nil
	}

	profile, err := _this.repo.FindByID(ctx, id)
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    http.StatusNotFound,
			Message: "not found",
			Result:  nil,
		}, err
	}

	payload := map[string]any{
		"id":         string(profile.GetID()),
		"user_id":    string(profile.GetUserID()),
		"full_name":  profile.GetFullName(),
		"avatar":     profile.GetAvatar(),
		"bio":        profile.GetBio(),
		"created_at": profile.GetCreatedAt(),
		"updated_at": profile.GetUpdatedAt(),
	}
	value, _ := structpb.NewValue(payload)
	return &profile_grpc.ProfileResponse{
		Code:    http.StatusAccepted,
		Message: "success",
		Result:  []*structpb.Value{value},
	}, nil
}
