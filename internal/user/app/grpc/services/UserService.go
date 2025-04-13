package grpc_services

import (
	"context"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
	pkgUtils "delivery/pkg/utils"

	"google.golang.org/protobuf/types/known/structpb"
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

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":    res.GetID(),
		"email": res.GetEmail(),
		"role":  res.GetRole(),
	})

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
		Result:  []*structpb.Value{stuctRes},
	}, nil
}

func (us *UserService) Delete(ctx context.Context, in *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {
	if ok, err := us.userRepo.Delete(ctx, in.ID); !ok || err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(struct {
		id string
	}{
		id: in.ID,
	})

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
		Result:  []*structpb.Value{stuctRes},
	}, nil

}

func (us *UserService) GetOne(ctx context.Context, in *user_grpc.GetUserRequest) (*user_grpc.Response, error) {

	res, err := us.userRepo.GetOne(ctx, in.GetId())
	if err != nil {
		return &user_grpc.Response{
			Code:    200,
			Message: "success",
			Result:  []*structpb.Value{},
		}, nil
	}

	resultMap := make([]*structpb.Value, 1)
	resultMap[0], _ = structpb.NewValue(map[string]any{
		"id":    res.GetID(),
		"email": res.GetEmail(),
		"role":  res.GetRole(),
	})

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  resultMap,
	}, nil
}

func (us *UserService) GetMany(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	if in.Offset <= 0 {
		in.Offset = 10
	}
	if in.Page <= 0 {
		in.Page = 1
	}

	res, err := us.userRepo.GetMany(ctx, in.Page, in.Offset)
	if err != nil {
		return &user_grpc.Response{
			Code:    200,
			Message: "success",
			Result:  []*structpb.Value{},
		}, nil
	}

	resultMap := make([]*structpb.Value, len(res))

	for i, r := range res {
		resultMap[i], _ = structpb.NewValue(map[string]any{
			"id":    r.GetID(),
			"email": r.GetEmail(),
			"role":  r.GetRole(),
		})
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  resultMap,
	}, nil
}

func (us *UserService) Search(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	if in.Offset <= 0 {
		in.Offset = 10
	}
	if in.Page <= 0 {
		in.Page = 1
	}

	res, err := us.userRepo.Search(ctx, in.Filter, in.Page, in.Offset)
	if err != nil {
		return &user_grpc.Response{
			Code:    200,
			Message: "success",
			Result:  []*structpb.Value{},
		}, nil
	}

	resultMap := make([]*structpb.Value, len(res))

	for i, r := range res {
		resultMap[i], _ = structpb.NewValue(map[string]any{
			"id":    r.GetID(),
			"email": r.GetEmail(),
			"role":  r.GetRole(),
		})
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  resultMap,
	}, nil
}

func (us *UserService) Update(ctx context.Context, in *user_grpc.UpdateUserRequest) (*user_grpc.Response, error) {
	u := &models.User{
		Email: in.Email,
		Role:  in.Role,
	}

	if in.Password != "" {
		u.Password = pkgUtils.Sha512(in.Password)
	}

	res, err := us.userRepo.Update(ctx, u)

	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":    res.GetID(),
		"email": res.GetEmail(),
		"role":  res.GetRole(),
	})

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
		Result:  []*structpb.Value{stuctRes},
	}, nil
}
