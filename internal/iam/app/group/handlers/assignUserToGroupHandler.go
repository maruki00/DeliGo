package groupHandlers

import (
	"context"
	pkgCqrs "deligo/pkg/cqrs"
)

type AssignUserToGroupHandler struct {
}

func (_this *AssignUserToGroupHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	return nil
}
