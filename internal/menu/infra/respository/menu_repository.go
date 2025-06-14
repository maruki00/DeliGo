package respositories

import (
	"context"

	shared_valueobject "github.com/maruki00/deligo/internal/shared/valueobject"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type MenuRepository struct {
	db *pkgPostgres.PGHandler
}

func NewMenuRepository(db *pkgPostgres.PGHandler) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}

func (_this *MenuRepository) Save(ctx context.Context) error {

	return nil
}

func (_this *MenuRepository) Update(ctx context.Context, id shared_valueobject.ID, menu any, products []shared_valueobject.ID) error {

	return nil
}

func (_this *MenuRepository) Delete(ctx context.Context, id shared_valueobject.ID) error {

	return nil
}

func (_this *MenuRepository) GetMenu(ctx context.Context, id shared_valueobject.ID) ([]any, error) {
	return nil, nil
}

func (_this *MenuRepository) SearchForMenu(ctx context.Context, query string) ([]any, error) {
	return nil, nil
}
