package grpc_services

import (
	"context"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/repositories"
	"encoding/json"
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

func (us *UserService) Create(ctx context.Context, createUserRequest *user_grpc.CreateUserRequest) (*user_grpc.UserResponse, error) {
	res, err := us.userRepo.Create(ctx, nil)
	if err != nil {
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	data, err := json.Marshal(res)
	if err != nil {
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}
	return &user_grpc.UserResponse{
		Code:    200,
		Message: "success",
		Result:  data,
	}, nil
}
func (us *UserService) Delete(context.Context, *user_grpc.DeleteUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) GetMany(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) GetOne(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) Search(context.Context, *user_grpc.EmptyUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) Update(context.Context, *user_grpc.UpdateUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
