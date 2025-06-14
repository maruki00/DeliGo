package queries

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetOneProfileQuery struct {
	ID sharedvo.ID
}

func (_this *GetOneProfileQuery) Name() string {
	return "GetOneProfileQuery"
}
