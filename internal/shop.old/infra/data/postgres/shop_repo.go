package postgres

import pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

type ShopRepo struct {
	db *pkgPostgres.PGHandler
}
