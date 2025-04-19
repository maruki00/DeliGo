package contracts

import (
	"context"
	"delivery/internal/iam/domain/entities"
)

type IProfileRepository interface {
	Create(ctx context.Context, entity entities.ProfileEntity) (entities.ProfileEntity, error)
	Delete(ctx context.Context, id string) (bool, error)
	Update(ctx context.Context, entity entities.ProfileEntity) (entities.ProfileEntity, error)
	GetOne(ctx context.Context, id string) (entities.ProfileEntity, error)
	GetMany(ctx context.Context, limit, offset int) ([]entities.ProfileEntity, error)
	Search(ctx context.Context, query string, limit, offset int) ([]entities.ProfileEntity, error)
}
