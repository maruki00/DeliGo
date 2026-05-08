package query

import sharedvo "github.com/maruki00/deligo/internal/shared/value_object"

type GetShopSquery struct {
	ID sharedvo.ID
}

func (_this *GetShopSquery) Name() string {
	return "GetShopSquery"
}
