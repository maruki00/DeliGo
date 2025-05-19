package grpc_services

import (
	"context"
	"deligo/internal/profile/app/profile/commands"
	profile_grpc "deligo/internal/profile/infra/grpc/profile"
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
	pkgCqrs "deligo/pkg/cqrs"
	"net/http"

	"google.golang.org/grpc"
)

type ProfileServerService struct {
	profile_grpc.UnimplementedProfileServiceServer
	CommandBus *pkgCqrs.CommandBus
	QueryBus   *pkgCqrs.QueryBus
}

func NewProfileService(cmdBus *pkgCqrs.CommandBus, qryBus *pkgCqrs.QueryBus) *ProfileServerService {
	return &ProfileServerService{
		CommandBus: cmdBus,
		QueryBus:   qryBus,
	}
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
	}, err
}

func (_this *ProfileServerService) Save(ctx context.Context, in *profile_grpc.CreateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {

	err := _this.CommandBus.Dispatch(ctx, &commands.SaveProfileCommand{
		ID:       shared_valueobject.NewID(),
		UserID:   shared_valueobject.ID(in.UserID),
		FullName: in.FullName,
		Avatar:   in.Avatar,
		Bio:      in.Bio,
	})
	return CommandCheck(err)

}
func (_this *ProfileServerService) Disable(ctx context.Context, in *profile_grpc.DisableProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {

	err := _this.CommandBus.Dispatch(ctx, &commands.DiscableProfileCommand{
		ID: shared_valueobject.ID(in.ID),
	})

	return CommandCheck(err)
}
func (_this *ProfileServerService) Update(ctx context.Context, in *profile_grpc.UpdateProfileRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {

	fields := make(map[string]any)
	for k, v := range in.Fields {
		fields[k] = v.(any)
	}
	err := _this.CommandBus.Dispatch(ctx, &commands.UpdateProfileCommand{
		ID:     shared_valueobject.ID(in.ID),
		Fields: fields,
	})
	return CommandCheck(err)
}
func (_this *ProfileServerService) UpdateAvatar(ctx context.Context, in *profile_grpc.UpdateProfileAvatareRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	err := _this.CommandBus.Dispatch(ctx, &commands.UpdateProfileAvatarCommand{
		ID:     shared_valueobject.ID(in.ID),
		Avatar: in.Avatar,
	})
	return CommandCheck(err)
}
func (_this *ProfileServerService) GetOne(ctx context.Context, in *profile_grpc.GETRequest, opts ...grpc.CallOption) (*profile_grpc.ProfileResponse, error) {
	return nil, nil
}
