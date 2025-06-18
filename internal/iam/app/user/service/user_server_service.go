package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/maruki00/deligo/internal/iam/domain/entity"
	"github.com/maruki00/deligo/internal/iam/infra/model"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
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

	command := &userCommand.CreateUserCommand{
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

	command := &userCommand.DeleteUserCommand{
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
	fields := in.GetFields()
	if fields == nil {
		return &user_grpc.Response{
			Code:    200,
			Message: "success",
			Details: nil,
		}, nil
	}

	var allowedField = map[string]bool{
		"email":    true,
		"username": true,
	}
	flds := make(map[string]interface{})

	for key, value := range fields {
		allowed, ok := allowedField[key]
		if !allowed || !ok {
			return &user_grpc.Response{
				Code:    400,
				Message: fmt.Sprintf("%s is not allowed to modify", key),
				Details: nil,
			}, nil
		}
		flds[key] = value
	}

	command := &userCommand.UpdateUserCommand{
		ID:     uuid.MustParse(in.ID),
		Fields: flds,
	}

	err := _this.commandBus.Dispatch(ctx, command)
	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: "Error : " + err.Error(),
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
	case "id":
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

	default:
		return &user_grpc.Response{
			Code:    400,
			Message: "invalid id",
			Details: nil,
		}, err

	}

	res, err := _this.queryBus.Dispatch(ctx, query)
	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Details: nil,
		}, err
	}
	entity := res.(entity.UserEntity)

	d, _ := structpb.NewStruct(map[string]any{
		"id":         string(entity.GetID()),
		"email":      entity.GetEmail(),
		"user_name":  entity.GetUsername(),
		"profile":    "",
		"updated_at": entity.GetUpdatedAt().String(),
		"created_at": entity.GetCreatedAt().String(),
	})

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: []*structpb.Struct{d},
	}, nil
}

func (_this *UserServerService) ListByTenant(ctx context.Context, in *user_grpc.GETRequest) (*user_grpc.Response, error) {

	params, err := pkgUtils.ParamsFromGrpc(ctx)
	if err != nil {
		return nil, err
	}
	tID := ""
	page := 1
	limit := 10
	if len(params["tenant_id"]) > 0 {
		tID = params["tenant_id"][0]
	}
	if len(params["page"]) > 0 {
		page, _ = strconv.Atoi(params["page]"][0])
	}

	if len(params["limit"]) > 0 {
		limit, _ = strconv.Atoi(params["limit"][0])
	}

	tenantID, _ := uuid.Parse(tID)

	query := &userQueries.ListUsersByTenantQuery{
		TenantID: tenantID,
		Pagination: shared_model.Pagination{
			Page:  page,
			Limit: limit,
		},
	}
	res, err := _this.queryBus.Dispatch(ctx, query)
	if err != nil {
		return nil, err
	}

	details := make([]*structpb.Struct, limit)
	index := 0
	for _, entity := range res.([]*model.User) {
		d, _ := structpb.NewStruct(map[string]any{
			"id":         string(entity.GetID()),
			"email":      entity.GetEmail(),
			"user_name":  entity.GetUsername(),
			"profile":    "",
			"updated_at": entity.GetUpdatedAt().String(),
			"created_at": entity.GetCreatedAt().String(),
		})
		details[index] = d
		index++
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Details: details[:index],
	}, nil
}
