package userServerServices

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	userQueries "deligo/internal/iam/app/user/queries"
	"deligo/internal/iam/domain/entities"
	user_grpc "deligo/internal/iam/infra/grpc/user"
	"deligo/internal/iam/infra/models"
	shared_models "deligo/internal/shared/infra/models"
	pkgCqrs "deligo/pkg/cqrs"
	pkgUtils "deligo/pkg/utils"
	"strconv"

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

	if err := in.Validate(); err != nil {
		return nil, err
	}

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
			Details: nil,
		}, err
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: nil,
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
			Details: nil,
		}, err
	}
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: nil,
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
			Details: nil,
		}, err
	}
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: nil,
	}, nil
}

func (_this *UserServerService) Find(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {

	params, err := pkgUtils.ParamsFromGrpc(ctx)
	if err != nil {
		return nil, err
	}
	value := ""
	filter := "value"
	if len(params["value"]) > 0 {
		value = params["value"][0]
	}
	if len(params["filter"]) > 0 {
		filter = params["filter"][0]

	}

	var query pkgCqrs.Query
	switch filter {
	case "username":
		query = &userQueries.FindUserByUsernameQuery{
			Key:      "username",
			Username: value,
		}
	case "email":
		query = &userQueries.FindUserByEmailQuery{
			Key:   "email",
			Email: value,
		}
	default:
		id, err := uuid.Parse(value)
		if err != nil {
			return &user_grpc.Response{
				Code:    400,
				Message: "invalid id",
				Details: nil,
			}, err
		}
		query = &userQueries.FindUserByIdQuery{
			Key: "id",
			ID:  id,
		}
	}

	res, err := _this.queryBus.Dispatch(ctx, query)
	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Details: nil,
		}, err
	}
	entity := res.(entities.UserEntity)
	var m = map[string]any{
		"id":         string(entity.GetID()),
		"email":      entity.GetEmail(),
		"user_name":  entity.GetUsername(),
		"profile":    "",
		"updated_at": entity.GetUpdatedAt().String(),
		"created_at": entity.GetCreatedAt().String(),
	}
	dd, _ := structpb.NewStruct(m)
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: dd,
	}, nil
}

func (_this *UserServerService) ListByTenant(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {

	params, err := pkgUtils.ParamsFromGrpc(ctx)
	if err != nil {
		return nil, err
	}
	tID := ""
	page := 1
	offset := 10
	if len(params["tenant_id"]) > 0 {
		tID = params["value"][0]
	}
	if len(params["page"]) > 0 {
		page, _ = strconv.Atoi(params["page]"][0])
	}

	if len(params["offset"]) > 0 {
		offset, _ = strconv.Atoi(params["offset"][0])
	}

	tenantID, _ := uuid.Parse(tID)

	query := &userQueries.ListUsersByTenantQuery{
		TenantID: tenantID,
		Pagination: shared_models.Pagination{
			Page:   page,
			Offset: offset,
		},
	}
	res, err := _this.queryBus.Dispatch(ctx, query)

	data := make([]any, offset)
	index := 0
	for _, entity := range res.([]*models.User) {
		data[index] = map[string]any{
			"id":         string(entity.GetID()),
			"email":      entity.GetEmail(),
			"user_name":  entity.GetUsername(),
			"profile":    "",
			"updated_at": entity.GetUpdatedAt().String(),
			"created_at": entity.GetCreatedAt().String(),
		}
		index++
	}

	x, _ := structpb.NewList(data)
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: x,
	}, nil
}
