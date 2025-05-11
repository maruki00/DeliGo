package queries

import "github.com/google/uuid"

type GetOneProfileQuery struct {
	ID uuid.UUID
}

func (_this *GetOneProfileQuery) Name() string {
	return "GetOneProfileQuery"
}
