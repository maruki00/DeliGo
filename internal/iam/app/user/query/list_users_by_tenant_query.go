package userQuery

import (
	"github.com/maruki00/deligo/pkg/pagination"

	"github.com/google/uuid"
)

type ListUsersByTenantQuery struct {
	TenantID   uuid.UUID
	Pagination pagination.Pagination
}

func (_this *ListUsersByTenantQuery) Name() string {
	return "ListUsersByTenantQuery"
}
