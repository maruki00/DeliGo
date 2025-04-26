package userServerServices

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	userQueries "deligo/internal/iam/app/user/queries"
	user_grpc "deligo/internal/iam/infra/grpc/user"
	pkgCqrs "deligo/pkg/cqrs"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
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

func (_this *UserServerService) Save(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {

	command := &userCommands.CreateUserCommand{
		ID:                uuid.New(),
		Username:          in.UserName,
		Email:             in.Email,
		Password:          in.Password,
		IsActive:          false,
		MFAEnabled:        false,
		PasswordChangedAt: nil,
		MFASecret:         "",
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

func (_this *UserServerService) Delete(ctx context.Context, in *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {

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

func (_this *UserServerService) Update(ctx context.Context, in *user_grpc.UpdateUserRequest) (*user_grpc.Response, error) {
	command := &userCommands.UpdateUserCommand{
		ID:     uuid.MustParse(in.ID),
		Fields: in.Fields,
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
	queryParams := in.GetQueryParams()
	var filterValue = "id"
	var valueValue = ""

	filter, ok := queryParams.Fields["filter"]
	if ok {
		filterValue = filter.GetStringValue()
	}

	value, ok := queryParams.Fields["value"]
	if ok {
		valueValue = value.GetStringValue()
	}
	var query pkgCqrs.Query
	switch filterValue {
	case "id":
		id, err := uuid.Parse(valueValue)
		if err != nil {
			return &user_grpc.Response{
				Code:    400,
				Message: "invalid id",
				Result:  nil,
			}, err
		}
		query = &userQueries.FindUserByIdQuery{
			ID: id,
		}
	case "username":
		query = &userQueries.FindUserByUsernameQuery{
			Username: valueValue,
		}
	case "email":
		query = &userQueries.FindUserByEmailQuery{
			Email: valueValue,
		}
	default:
		return &user_grpc.Response{
			Code:    400,
			Message: "invalid filter",
			Result:  nil,
		}, fmt.Errorf("invalid filter")
	}
	res, err := _this.queryBus.Dispatch(ctx, query)
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
		Result:  res.([]*structpb.Value),
	}, nil
}

func (_this *UserServerService) ListByTenant(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {

	queryParams := in.GetQueryParams()
	var tenantID = ""

	filter, ok := queryParams.Fields["tenantId"]
	if ok {
		tenantID = filter.GetStringValue()

	}

	id, _ := uuid.Parse(tenantID)

	query := &userQueries.ListUsersByTenantQuery{
		TenantID: id,
	}
	res, err := _this.queryBus.Dispatch(ctx, query)
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
		Result:  res.([]*structpb.Value),
	}, nil
}
