package query

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetProductsByIDsQuery struct {
	IDs    []*sharedvo.ID
	Page   int
	Offest int
}

func (_this *GetProductsByIDsQuery) Name() string {
	return "GetProductsByIDsQuery"
}
