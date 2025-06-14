package query

import sharedvo "deligo/internal/shared/valueobject"

type GetProductByMultipleIDQuery struct {
	IDs    []*sharedvo.ID
	Page   int
	Offest int
}

func (_this *GetProductByMultipleIDQuery) Name() string {
	return "GetProductByMultipleIDQuery"
}
