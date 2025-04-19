package pkgPostgres

import (
	"log/slog"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PGHandler struct {
	DB       *gorm.DB
	MaxTries int
	Timeout  int
}

func (pg *PGHandler) SetDB(db *gorm.DB) {
	pg.DB = db
}

func NewDB(dsn string) (*PGHandler, error) {

	objDB := &PGHandler{
		DB:       nil,
		MaxTries: MAX_TRIES,
		Timeout:  TIMEOUT,
	}
	var err error
	for i := range objDB.MaxTries {
		slog.Info("trying to connect to Postgres.", i, " times")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			continue
		}
		objDB.DB = db
		return objDB, nil
	}

	return nil, err
}

func (pg *PGHandler) GetDB() *gorm.DB {
	return pg.DB
}

func (pg *PGHandler) Close() {

}
