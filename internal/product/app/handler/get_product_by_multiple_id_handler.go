package handler

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type GetProductByMultipleIDHandler struct {
}

func (_this *GetProductByMultipleIDHandler) handler(ctx context.Context, q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
