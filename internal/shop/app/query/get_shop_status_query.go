package query

import sharedvo "deligo/internal/shared/valueobject"

type GetShopStatusQuery struct {
	ID sharedvo.ID
}

func (_this *GetShopStatusQuery) Name() string {
	return "GetShopSquery"
}
