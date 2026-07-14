package repository

import (
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type AuthzRepository struct {
	db *pkgPostgres.PGHandler
}

func NewAuthzRepository(db *pkgPostgres.PGHandler) *AuthzRepository {
	return &AuthzRepository{
		db: db,
	}
}
