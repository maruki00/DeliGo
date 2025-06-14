package queries

import sharedvo "deligo/internal/shared/valueobject"

type GetOneProfileQuery struct {
	ID sharedvo.ID
}

func (_this *GetOneProfileQuery) Name() string {
	return "GetOneProfileQuery"
}
