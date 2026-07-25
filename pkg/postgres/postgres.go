package pkgPostgres

import (
	"fmt"
	"log"
	"time"

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
	for i := 1; i <= objDB.MaxTries; i++ {
		objDB.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			log.Println("Successfully connected to Postgres!")
			return objDB, nil
		}

		log.Printf("Postgres not ready yet (attempt %d/%d): %v. Retrying in 2s...", i, objDB.MaxTries, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to postgres after %d attempts: %w", objDB.MaxTries, err)

	// for i := range objDB.MaxTries {
	// 	slog.Info("trying to connect to Postgres.", fmt.Sprintf("%d", i), " times")
	// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 		Logger: logger.Default.LogMode(logger.Info),
	// 	})
	// 	if err != nil {
	// 		slog.Error("failed", "connect to db", err.Error())
	// 		continue
	// 	}
	// 	objDB.DB = db
	// 	return objDB, nil
	// }
	// return nil, err
}

func (pg *PGHandler) GetDB() *gorm.DB {
	return pg.DB
}

func (pg *PGHandler) Close() {
	db, err := pg.GetDB().DB()
	if err != nil {
		return
	}
	db.Close()
}
