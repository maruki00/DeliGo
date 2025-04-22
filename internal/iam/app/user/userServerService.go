package userServerServices

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	user_grpc "deligo/internal/iam/infra/grpc/user"
	pkgCqrs "deligo/pkg/cqrs"

	"github.com/google/uuid"
)

type UserServerService struct {
	user_grpc.UnimplementedUserServiceServer
	commandBus *pkgCqrs.CommandBus
	queryBus   *pkgCqrs.QueryBus
}

func NewUserUseCase(
	commandBus *pkgCqrs.CommandBus,
	queryBus *pkgCqrs.QueryBus) *UserServerService {
	return &UserServerService{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}


// func (UnimplementedUserServiceServer) Find(context.Context, *GETRequest) (*Response, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
// }
// func (UnimplementedUserServiceServer) ListByTenant(context.Context, *GETRequest) (*Response, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method ListByTenant not implemented")
// }



func (_this *UserServerService) Save(ctx context.Context, in *CreateUserRequest) (*user_grpc.Response, error) {
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

func (_this *UserServerService) Delete(context.Context, *DeleteUserRequest) (*user_grpc.Response, error) {

	command := &userCommands.DeleteUserCommand{
		ID: uuid.MustParse(in.ID),
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

func (_this *UserServerService) Update(ctx context.Context, in *UpdateUserRequest) (*user_grpc.Response, error)  {
	command := &userCommands.UpdateUserCommand{
		ID:         uuid.MustParse(in.ID),
		Fields:    in.Fields,
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


func (_this *UserServerService) Find(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {
	in.QueryParams = map[string]string{
		"filter": in.Filter,
		"email":    in.Email,
	}
	query := &userQueries.
	return nil, nil
}



func (_this *UserServerService) ListByTenant(ctx context.Context, in *GETRequest) (*user_grpc.Response, error) {
	query := &userQueries.ListUsersByTenantQuery{
		TenantID: uuid.MustParse(in.TenantId),
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	err := _this.queryBus.Dispatch(ctx, query)
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

