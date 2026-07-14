package repository

import (
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type UserRepository struct {
	db *pkgPostgres.PGHandler
}

func NewUserRepository(db *pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
