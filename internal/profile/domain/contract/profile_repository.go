package contract

import (
	"context"
	"github.com/maruki00/deligo/internal/profile/domain/entity"
	"github.com/maruki00/deligo/internal/profile/infra/model"
)

type IPorofileRepository interface {
	Save(ctx context.Context, entity entity.ProfileEntity) error
	Disable(ctx context.Context, id string) error
	FindByUserID(ctx context.Context, id string) (*model.Profile, error)
	Update(context.Context, string, map[string]any) error
	UpdateAvatar(ctx context.Context, id string, avatar string) error
}
