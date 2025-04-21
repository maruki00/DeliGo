package groupHandlers

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type CreateGroupHandler struct {
}

func (_this *CreateGroupHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	return nil
}
