package services

import (
	"context"
	"delivery/internal/product/domian/dtos"
	"delivery/internal/product/domian/ports"
	"delivery/internal/product/domian/repositories"
	"delivery/internal/product/infrastructure/models"
	shared_contracts "delivery/internal/shared/domain/contracts"
	shared_models "delivery/internal/shared/infra/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	productRepository repositories.IProductRepository
	outputPort        ports.ProductOutputPort
}

func NewProductService(productRepository repositories.IProductRepository, outputPort ports.ProductOutputPort) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		outputPort:        outputPort,
	}
}

func (obj *ProductService) Insert(ctx context.Context, dto *dtos.InsertProductDTO) shared_contracts.ViewModel {

	res, err := obj.productRepository.Insert(ctx, &models.Product{
		Label: dto.Label,
		Type:  dto.Type,
		Price: dto.Price,
	})
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) Search(ctx context.Context, dto *dtos.SearchProductDTO) shared_contracts.ViewModel {

	res, err := obj.productRepository.Search(ctx, dto.Query)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"products": res},
	})
}

func (obj *ProductService) Update(ctx context.Context, dto *dtos.UpdateProductDTO) shared_contracts.ViewModel {

	res, err := obj.productRepository.Update(ctx, dto.Id, map[string]interface{}{
		"id":    dto.Id,
		"label": dto.Label,
		"type":  dto.Type,
		"price": dto.Price,
	})

	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) Delete(ctx context.Context, dto *dtos.DeleteProductDTO) shared_contracts.ViewModel {

	res, err := obj.productRepository.Delete(ctx, dto.Id)

	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not delete the product : " + err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) GetById(ctx context.Context, dto *dtos.GetProductByIdDTO) shared_contracts.ViewModel {

	res, err := obj.productRepository.GetById(ctx, dto.Id)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not get the product or doesnt exsits",
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) MultipleProducts(ctx context.Context, dto *dtos.MultipleProductsDTO) shared_contracts.ViewModel {

	ids := []int{}
	strIds := strings.Split(dto.Ids, ",")

	for _, id := range strIds {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, int(i))
	}

	res, err := obj.productRepository.GetProductByMultipleId(ctx, ids)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not get the product or doesnt exsits",
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"products": res},
	})
}
