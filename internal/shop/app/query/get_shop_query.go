package query

import sharedvo "deligo/internal/shared/valueobject"

type GetShopSquery struct {
	ID sharedvo.ID
}

func (_this *GetShopSquery) Name() string {
	return "GetShopSquery"
}
