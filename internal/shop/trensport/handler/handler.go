package handler

import (
	"github.com/maruki00/deligo/internal/shop/app"
	menuDomain "github.com/maruki00/deligo/internal/shop/domain/menu"
	productDomain "github.com/maruki00/deligo/internal/shop/domain/product"
	shopDomain "github.com/maruki00/deligo/internal/shop/domain/shop"
)

type Handler struct {
	app            *app.App
	shopService    shopDomain.Service
	menuService    menuDomain.Service
	productService productDomain.Service
}

func New(instance *app.App) *Handler {
	return &Handler{
		app:            instance,
		shopService:    instance.ShopService,
		menuService:    instance.MenuService,
		productService: instance.ProductService,
	}
}
