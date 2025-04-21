package groupHandlers

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type DeleteGroupHandler struct {
}

func (_this *DeleteGroupHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	return nil
}
