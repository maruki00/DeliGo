package controllers

import (
	"context"
	"delivery/internal/product/app/services"
	"delivery/internal/product/domian/dtos"
	"delivery/internal/product/domian/ports"
	"delivery/internal/product/infrastructure/models"
	"delivery/internal/product/infrastructure/repositories"
	"delivery/internal/product/userGetway/adapters/presenters"
	"delivery/internal/product/userGetway/requests"
	shared_core "delivery/internal/shared/infra/core"
	"delivery/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductController struct {
	Validate *validator.Validate
	inPort   ports.ProductInputPort
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		inPort: services.NewProductService(
			repositories.NewProductRepository(db, &models.Product{}),
			&presenters.ProductPresenter{},
		),
		Validate: validator.New(),
	}
}

func (obj *ProductController) Insert(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Insert(c, &dtos.InsertProductDTO{
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Search(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &requests.SearchProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Search(c, &dtos.SearchProductDTO{
		Query: request.Query,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Update(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &requests.UpdateProductRequerst{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", "validation : "+err.Error())
		return
	}
	res := obj.inPort.Update(c, &dtos.UpdateProductDTO{
		Id:    request.Id,
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Delete(ctx *gin.Context) {

	// c, Cancel := context.WithCancel(context.Background())
	c, Cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	defer Cancel()

	request := &requests.DeleteProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Delete(c, &dtos.DeleteProductDTO{
		Id: request.Id,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) GetProduct(ctx *gin.Context) {
	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()

	request := &requests.GetProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.GetById(c, &dtos.GetProductByIdDTO{
		Id: request.Id,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) MultipleProducts(ctx *gin.Context) {
	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &requests.MultipleProductstRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.MultipleProducts(c, &dtos.MultipleProductsDTO{
		Ids: request.Ids,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}
