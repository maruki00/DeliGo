package handler

import (
	"github.com/maruki00/deligo/internal/product/domian/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
)

type SaveProductHandler struct {
	repo contract.IProductRepository
}

func NewSaveProductHandler(repo contract.IProductRepository) SaveProductHandler {
	return SaveProductHandler{
		repo: repo,
	}
}

func (_this *SaveProductHandler) handler(c pkgCqrs.Command) error {

	return nil
}
