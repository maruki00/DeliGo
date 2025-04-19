package contracts

import (
	"context"
	"delivery/internal/user/domain/entities"
	"delivery/internal/user/infra/models"
)

type IUserRepository interface {
	Create(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error)
	Delete(ctx context.Context, id string) (bool, error)
	Update(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error)
	GetOne(ctx context.Context, userId string) (entities.UserEntity, error)
	GetMany(ctx context.Context, offset, limit int32) ([]*models.User, error)
	Search(ctx context.Context, query string, offset, limit int32) ([]*models.User, error)
}
