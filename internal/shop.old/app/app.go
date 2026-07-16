package app

import (
	"fmt"

	"github.com/maruki00/deligo/internal/shop/app/service"
	shopDomain "github.com/maruki00/deligo/internal/shop/domain/shop"
	menuDomain "github.com/maruki00/deligo/internal/shop/domain/menu"
	productDomain "github.com/maruki00/deligo/internal/shop/domain/product"
	postgresRepo "github.com/maruki00/deligo/internal/shop/infra/data/postgres"
)

type App struct {
	ShopService  shopDomain.Service
	MenuService  menuDomain.Service
	ProductService productDomain.Service
}

func Init() (*App, func(), error) {
	repo := postgresRepo.NewRepository()
	app := &App{
		ShopService:  service.NewShopService(repo.Shop),
		MenuService:  service.NewMenuService(repo.Menu),
		ProductService: service.NewProductService(repo.Product),
	}
	return app, func() {}, nil
}

func (a *App) Health() string {
	return fmt.Sprintf("shop-service-up")
}


