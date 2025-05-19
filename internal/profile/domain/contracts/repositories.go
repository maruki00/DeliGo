package contracts

import (
	"context"
	"deligo/internal/profile/domain/entities"
	"deligo/internal/profile/infra/models"
)

type IPorofileRepository interface {
	Save(ctx context.Context, entity entities.ProfileEntity) error
	Disable(ctx context.Context, id string) error
	FindByUserID(ctx context.Context, id string) (*models.Profile, error)
	Update(context.Context, string, map[string]any) error
	UpdateAvatar(ctx context.Context, id string, avatar string) error
}
