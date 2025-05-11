package contracts

import (
	"context"
	"deligo/internal/profile/infra/models"
)

type IPorofileRepository interface {
	Save(context.Context, *models.Profile) error
	Disable(context.Context, *models.Profile) error
	FindByUserID(context.Context, string) (*models.Profile, error)
	Update(context.Context, string, map[string]any) error
	UpdateAvatar(context.Context, string, string) error
}
