package handler

import pkgCqrs "deligo/pkg/cqrs"

type DeleteProductHandler struct {
}

func (_this *DeleteProductHandler) handler(ctx, c pkgCqrs.Command) error {

	return nil
}
