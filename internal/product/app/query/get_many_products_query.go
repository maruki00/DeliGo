package query

import sharedvo "deligo/internal/shared/valueobject"

type GetManyProductQuery struct {
	IDs    []*sharedvo.ID
	Page   int
	Offest int
}

func (_this *GetManyProductQuery) Name() string {
	return "GetManyProductQuery"
}
