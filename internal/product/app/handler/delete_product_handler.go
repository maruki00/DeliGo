package handler

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type DeleteProductHandler struct {
}

func (_this *DeleteProductHandler) handler(ctx context.Context, c pkgCqrs.Command) error {

	return nil
}
