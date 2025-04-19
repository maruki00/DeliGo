package pkgPostgres

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormHandler struct {
	DB       *gorm.DB
	MaxTries int
	Timeout  int
}

func NewGormHandler(dsn string) (*GormHandler, error) {
	gormObj := &GormHandler{
		DB:       nil,
		MaxTries: MAX_TRIES,
		Timeout:  TIMEOUT,
	}
	var err error
	for i := range gormObj.MaxTries {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			continue
		}
		gormObj.DB = db
		return gormObj, nil
	}

	return nil, err

}

func (pg *GormHandler) SetDB(db *gorm.DB) {
	pg.DB = db
}

func (pg *GormHandler) GetDB() *gorm.DB {
	return pg.DB
}

func (pg *GormHandler) Close() {
	_ = pg.GetDB()
}
