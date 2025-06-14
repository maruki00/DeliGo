package handler

import (
	"deligo/internal/product/domian/contract"
	pkgCqrs "deligo/pkg/cqrs"
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
