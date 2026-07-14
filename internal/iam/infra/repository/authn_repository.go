package repository

import (
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type AuthnRepository struct {
	db *pkgPostgres.PGHandler
}

func NewAuthnRepository(db *pkgPostgres.PGHandler) *AuthnRepository {
	return &AuthnRepository{
		db: db,
	}
}
