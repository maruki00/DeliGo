package grpc_services

import (
	"context"
	"delivery/internal/user/domain/entities"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
)

type UserService struct {
	user_grpc.UnimplementedUserServiceServer
	userRepo *repositories.UserRepository
}

var _ entities.UserEntity = (*models.User)(nil)

func (us *UserService) Create(ctx context.Context, createUserRequest *user_grpc.CreateUserRequest) (*user_grpc.UserResponse, error) {
	res, err := us.userRepo.Create(ctx, nil)
	if err != nil {
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	if err != nil {
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}
	// anypb.UnmarshalTo(res, anyRes)
	return &user_grpc.UserResponse{
		Code:    200,
		Message: "success",
		Result:  []byte(res),
	}, nil
}
func (us *UserService) Delete(context.Context, *user_grpc.DeleteUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) GetMany(context.Context, *user_grpc.EmptyUserResponse) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) GetOne(context.Context, *user_grpc.EmptyUserResponse) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) Search(context.Context, *user_grpc.EmptyUserResponse) (*user_grpc.UserResponse, error) {
	return nil, nil
}
func (us *UserService) Update(context.Context, *user_grpc.UpdateUserRequest) (*user_grpc.UserResponse, error) {
	return nil, nil
}
