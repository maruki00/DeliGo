package postgres

import (
	"fmt"
	"log"

	menuDomain "github.com/maruki00/deligo/internal/shop/domain/menu"
	productDomain "github.com/maruki00/deligo/internal/shop/domain/product"
	shopDomain "github.com/maruki00/deligo/internal/shop/domain/shop"
	shopModel "github.com/maruki00/deligo/internal/shop/infra/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepositoryBundle struct {
	Shop    shopDomain.Repository
	Menu    menuDomain.Repository
	Product productDomain.Repository
}

func NewRepository() *RepositoryBundle {
	dsn := "host=localhost user=postgres password=postgres dbname=deligo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("postgres unavailable, falling back to sqlite: %v", err)
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("failed to initialize fallback database: %v", err))
		}
	}

	_ = db.AutoMigrate(&shopModel.Shop{}, &shopModel.Menu{}, &shopModel.Product{})

	return &RepositoryBundle{
		Shop:    NewShopRepository(db),
		Menu:    NewMenuRepository(db),
		Product: NewProductRepository(db),
	}
}
