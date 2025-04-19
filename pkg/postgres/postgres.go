package pkgPostgres

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDB interface {
	*sql.DB | *gorm.DB
}

type SQLDB interface {
	*sql.DB
}

type GORMDB interface {
	*gorm.DB
}

type PGHandler[T PGDB] struct {
	DB       T
	MaxTries int
	Timeout  int
}

func (pg *PGHandler[T]) SetDB(db T) {
	pg.DB = db
}

func NewDB[T PGDB](dsn string) (T, error) {
	var zero T

	switch any(zero).(type) {
	case *sql.DB:
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return zero, err
		}
		return any(db).(T), nil
	case *gorm.DB:
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return zero, err
		}
		return any(db).(T), nil
	default:
		return zero, errors.New("database type is not supported")
	}
}

func (pg *PGHandler[T]) GetDB() T {
	return pg.DB
}

func (pg *PGHandler[T]) Close() {
	pg.Close()
}
