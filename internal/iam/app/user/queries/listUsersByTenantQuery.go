package userQueries

import (
	shared_models "deligo/internal/shared/infra/models"

	"github.com/google/uuid"
)

type ListUsersByTenantQuery struct {
	TenantID   uuid.UUID
	Pagination shared_models.Pagination
}

func (_this *ListUsersByTenantQuery) Name() string {
	return "ListUsersByTenantQuery"
}
