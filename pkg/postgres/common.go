package pkgPostgres

import "gorm.io/gorm"

const (
	MAX_TRIES = 3
	TIMEOUT   = 1
)

type DBHandler interface {
	SetDB(db *gorm.DB)
	GetDB() *gorm.DB
	Close()
}
