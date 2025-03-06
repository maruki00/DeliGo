package repositories

import (
	"errors"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	sync.Mutex
	db *gorm.DB
}

func NewCustomerRepository(dsn string) (*CustomerRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, errors.New("could not connect to server : " + err.Error())
	}
	return &CustomerRepository{
		db: db,
	}, nil
}
