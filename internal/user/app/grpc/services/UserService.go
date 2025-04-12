package grpc_services

import (
	"context"
	"delivery/internal/user/domain/entities"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
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

func (us *UserService) Create(ctx context.Context, createUserRequest *user_grpc.CreateUserRequest) (*user_grpc.UserResponse, error) {
	res, err := us.userRepo.Create(ctx, &models.User{
		ID:       "12345",
		Email:    "12345",
		Password: "12345",
		Role:     "12345",
	})

	if err != nil {
		fmt.Println("25: ", err.Error())
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	data, err := anypb.New(&user_grpc.User{
		// state:    "",
		ID:       res.GetID(),
		Email:    res.GetEmail(),
		Password: res.GetPassword(),
		Role:     res.GetRole(),
	})
	if err != nil {
		return &user_grpc.UserResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	// data, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println("35: ", err.Error())
	// 	return &user_grpc.UserResponse{
	// 		Code:    400,
	// 		Message: err.Error(),
	// 		Result:  nil,
	// 	}, err
	// }

	fmt.Println(data)
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
