package query

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetManyProductsQuery struct {
	IDs    []*sharedvo.ID
	Page   int
	Offest int
}

func (_this *GetManyProductsQuery) Name() string {
	return "GetManyProductsQuery"
}
