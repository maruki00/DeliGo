package userQueries

import "github.com/google/uuid"

type ListUsersByTenantQuery struct {
	TenantID uuid.UUID
}

func (_this *ListUsersByTenantQuery) QueryName() string {
	return "ListUsersByTenantQuery"
}
