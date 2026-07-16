package postgres

import (
	"context"

	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type MenuRepo struct {
	db *pkgPostgres.PGHandler
}

func NewMenuRepository(db *pkgPostgres.PGHandler) *MenuRepo {
	return &MenuRepo{
		db: db,
	}
}

func (_this *MenuRepo) Save(ctx context.Context) error {

	return nil
}

func (_this *MenuRepo) Update(ctx context.Context, id shared_valueobject.ID, menu any, products []shared_valueobject.ID) error {

	return nil
}

func (_this *MenuRepo) Delete(ctx context.Context, id shared_valueobject.ID) error {

	return nil
}

func (_this *MenuRepo) GetMenu(ctx context.Context, id shared_valueobject.ID) ([]any, error) {
	return nil, nil
}

func (_this *MenuRepo) SearchForMenu(ctx context.Context, query string) ([]any, error) {
	return nil, nil
}
