package query

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type GetShopSquery struct {
	ID shared_valueobject.ID
}

func (_this *GetShopSquery) Name() string {
	return "GetShopSquery"
}
