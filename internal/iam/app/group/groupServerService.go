package GroupServerServices

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	user_grpc "deligo/internal/iam/infra/grpc/user"
	pkgCqrs "deligo/pkg/cqrs"

	"github.com/google/uuid"
)

type GroupServerService struct {
	user_grpc.UnimplementedUserServiceServer
	commandBus *pkgCqrs.CommandBus
	queryBus   *pkgCqrs.QueryBus
}

func NewUserUseCase(
	commandBus *pkgCqrs.CommandBus,
	queryBus *pkgCqrs.QueryBus) *GroupServerService {
	return &GroupServerService{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

func (_this *GroupServerService) Create(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {
	command := &userCommands.CreateUserCommand{
		ID:         uuid.New(),
		Username:   in.UserName,
		Email:      in.Email,
		Password:   in.Password,
		IsActive:   false,
		MFAEnabled: false,
	}
	err := _this.commandBus.Dispatch(ctx, command)
	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  nil,
	}, nil
}

func (_this *GroupServerService) Delete(ctx context.Context, in *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (_this *GroupServerService) AssignUserToGroup(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (_this *GroupServerService) UnassignUserToGroup(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (_this *GroupServerService) GetUserGroups(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
