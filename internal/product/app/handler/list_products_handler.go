package handler

import "deligo/internal/product/domian/contract"

type ListProductHandler struct {
	repo contract.IProductRepository
}

func NewListProductHandler(repo contract.IProductRepository) ListProductHandler {
	return ListProductHandler{
		repo: repo,
	}
}

func (_this *ListProductHandler) Handler() (interface{}, error) {

	return nil, nil
}
