package query

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetProductByMultipleIDQuery struct {
	IDs    []*sharedvo.ID
	Page   int
	Offest int
}

func (_this *GetProductByMultipleIDQuery) Name() string {
	return "GetProductByMultipleIDQuery"
}
