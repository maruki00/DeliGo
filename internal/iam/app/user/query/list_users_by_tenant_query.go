package userQuery

import (
	shared_model "github.com/maruki00/deligo/internal/shared/model"

	"github.com/google/uuid"
)

type ListUsersByTenantQuery struct {
	TenantID   uuid.UUID
	Pagination shared_model.Pagination
}

func (_this *ListUsersByTenantQuery) Name() string {
	return "ListUsersByTenantQuery"
}
