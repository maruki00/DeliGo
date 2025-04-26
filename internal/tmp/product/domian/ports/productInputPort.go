package ports

import (
	"context"
	"deligo/internal/product/domian/dtos"
	shared_contracts "deligo/internal/shared/domain/contracts"
)

type ProductInputPort interface {
	Insert(ctx context.Context, dto *dtos.InsertProductDTO) shared_contracts.ViewModel
	Search(ctx context.Context, dto *dtos.SearchProductDTO) shared_contracts.ViewModel
	Update(ctx context.Context, dto *dtos.UpdateProductDTO) shared_contracts.ViewModel
	Delete(ctx context.Context, dto *dtos.DeleteProductDTO) shared_contracts.ViewModel
	GetById(ctx context.Context, dto *dtos.GetProductByIdDTO) shared_contracts.ViewModel
	MultipleProducts(ctx context.Context, dto *dtos.MultipleProductsDTO) shared_contracts.ViewModel
}
