package postgres

import (
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type AuthnRepo struct {
	db *pkgPostgres.PGHandler
}

func NewAuthnRepo(db *pkgPostgres.PGHandler) *AuthnRepo {
	return &AuthnRepo{
		db: db,
	}
}
