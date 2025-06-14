package handler

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetManyProductsHandler struct {
}

func (_this *GetManyProductsHandler) handler(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
