package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/shop/app"
	"github.com/maruki00/deligo/internal/shop/trensport/handler"
)

func NewRouter(instance *app.App) *gin.Engine {
	r := gin.Default()
	h := handler.New(instance)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	shops := r.Group("/shops")
	shops.POST("", h.CreateShop)
	shops.GET("", h.ListShops)
	shops.GET(":id", h.GetShop)
	shops.PUT(":id", h.UpdateShop)
	shops.DELETE(":id", h.DeleteShop)
	shops.POST(":id/menus", h.CreateMenu)
	shops.GET(":id/menus", h.ListMenusByShop)
	shops.POST(":id/products", h.CreateProduct)
	shops.GET(":id/products", h.ListProductsByShop)

	menus := r.Group("/menus")
	menus.GET(":id", h.GetMenu)
	menus.DELETE(":id", h.DeleteMenu)
	menus.POST(":id/items", h.AddMenuItem)
	menus.DELETE(":id/items/:productID", h.RemoveMenuItem)

	products := r.Group("/products")
	products.GET(":id", h.GetProduct)
	products.PUT(":id", h.UpdateProduct)
	products.PATCH(":id/stock", h.AdjustStock)
	products.DELETE(":id", h.DeleteProduct)

	return r
}
