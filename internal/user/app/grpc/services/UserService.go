package grpc_services

import (
	"context"
	"delivery/internal/user/domain/entities"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
	pkgUtils "delivery/pkg/utils"
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"
)

type UserService struct {
	user_grpc.UnimplementedUserServiceServer
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

var _ entities.UserEntity = (*models.User)(nil)

func (us *UserService) Create(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {
	res, err := us.userRepo.Create(ctx, &models.User{
		Email:    in.Email,
		Password: pkgUtils.Sha512(in.Password),
		Role:     in.Role,
	})

	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	data, err := anypb.New(&user_grpc.User{
		ID:       res.GetID(),
		Email:    res.GetEmail(),
		Password: res.GetPassword(),
		Role:     res.GetRole(),
	})
	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	fmt.Println(data)
	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  data,
	}, nil
}
func (us *UserService) Delete(context.Context, *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) GetMany(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) GetOne(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) Search(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) Update(context.Context, *user_grpc.UpdateUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
