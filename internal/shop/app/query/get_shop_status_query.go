package query

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetShopStatusQuery struct {
	ID sharedvo.ID
}

func (_this *GetShopStatusQuery) Name() string {
	return "GetShopSquery"
}
