package usecases

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	user_grpc "deligo/internal/iam/infra/grpc/user"
	pkgCqrs "deligo/pkg/cqrs"

	"github.com/google/uuid"
)

type UserUseCase struct {
	user_grpc.UnimplementedUserServiceServer
	commandBus *pkgCqrs.CommandBus
	queryBus   *pkgCqrs.QueryBus
}

func NewUserUseCase(
	commandBus *pkgCqrs.CommandBus,
	queryBus *pkgCqrs.QueryBus) *UserUseCase {
	return &UserUseCase{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

func (_this *UserUseCase) Create(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {
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

// func (us *UserUseCase) Delete(ctx context.Context, in *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {
// 	if ok, err := us.userRepo.Delete(ctx, in.ID); !ok || err != nil {
// 		return &user_grpc.Response{
// 			Code:    400,
// 			Message: err.Error(),
// 			Result:  nil,
// 		}, err
// 	}

// 	stuctRes, err := structpb.NewValue(struct {
// 		id string
// 	}{
// 		id: in.ID,
// 	})

// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    400,
// 			Message: err.Error(),
// 			Result:  nil,
// 		}, err
// 	}

// 	return &user_grpc.Response{
// 		Code:    200,
// 		Message: "success",
// 		Result:  []*structpb.Value{stuctRes},
// 	}, nil

// }

// func (us *UserUseCase) GetOne(ctx context.Context, in *user_grpc.GetUserRequest) (*user_grpc.Response, error) {

// 	res, err := us.userRepo.GetOne(ctx, in.GetId())
// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    200,
// 			Message: "success",
// 			Result:  []*structpb.Value{},
// 		}, nil
// 	}

// 	resultMap := make([]*structpb.Value, 1)
// 	resultMap[0], _ = structpb.NewValue(map[string]any{
// 		"id":    res.GetID(),
// 		"email": res.GetEmail(),
// 		"role":  res.GetRole(),
// 	})

// 	return &user_grpc.Response{
// 		Code:    200,
// 		Message: "success",
// 		Result:  resultMap,
// 	}, nil
// }

// func (us *UserUseCase) GetMany(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
// 	if in.Offset <= 0 {
// 		in.Offset = 10
// 	}
// 	if in.Page <= 0 {
// 		in.Page = 1
// 	}

// 	res, err := us.userRepo.GetMany(ctx, in.Page, in.Offset)
// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    200,
// 			Message: "success",
// 			Result:  []*structpb.Value{},
// 		}, nil
// 	}

// 	resultMap := make([]*structpb.Value, len(res))

// 	for i, r := range res {
// 		resultMap[i], _ = structpb.NewValue(map[string]any{
// 			"id":    r.GetID(),
// 			"email": r.GetEmail(),
// 			"role":  r.GetRole(),
// 		})
// 	}

// 	return &user_grpc.Response{
// 		Code:    200,
// 		Message: "success",
// 		Result:  resultMap,
// 	}, nil
// }

// func (us *UserUseCase) Search(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
// 	if in.Offset <= 0 {
// 		in.Offset = 10
// 	}
// 	if in.Page <= 0 {
// 		in.Page = 1
// 	}

// 	res, err := us.userRepo.Search(ctx, in.Filter, in.Page, in.Offset)
// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    200,
// 			Message: "success",
// 			Result:  []*structpb.Value{},
// 		}, nil
// 	}

// 	resultMap := make([]*structpb.Value, len(res))

// 	for i, r := range res {
// 		resultMap[i], _ = structpb.NewValue(map[string]any{
// 			"id":    r.GetID(),
// 			"email": r.GetEmail(),
// 			"role":  r.GetRole(),
// 		})
// 	}

// 	return &user_grpc.Response{
// 		Code:    200,
// 		Message: "success",
// 		Result:  resultMap,
// 	}, nil
// }

// func (us *UserUseCase) Update(ctx context.Context, in *user_grpc.UpdateUserRequest) (*user_grpc.Response, error) {
// 	u := &models.User{
// 		Email: in.Email,
// 		Role:  in.Role,
// 	}

// 	if in.Password != "" {
// 		u.Password = pkgUtils.Sha512(in.Password)
// 	}

// 	res, err := us.userRepo.Update(ctx, u)

// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    400,
// 			Message: err.Error(),
// 			Result:  nil,
// 		}, err
// 	}

// 	stuctRes, err := structpb.NewValue(map[string]any{
// 		"id":    res.GetID(),
// 		"email": res.GetEmail(),
// 		"role":  res.GetRole(),
// 	})

// 	if err != nil {
// 		return &user_grpc.Response{
// 			Code:    400,
// 			Message: err.Error(),
// 			Result:  nil,
// 		}, err
// 	}

// 	return &user_grpc.Response{
// 		Code:    200,
// 		Message: "success",
// 		Result:  []*structpb.Value{stuctRes},
// 	}, nil
// }
