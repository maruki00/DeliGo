package queries

import (
	shared_valueobject "deligo/internal/shared/domain/valueObjects"
)

type GetOneProfileQuery struct {
	ID shared_valueobject.ID
}

func (_this *GetOneProfileQuery) Name() string {
	return "GetOneProfileQuery"
}
