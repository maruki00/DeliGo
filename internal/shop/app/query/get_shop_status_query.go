package queries

import shared_valueobject "deligo/internal/shared/domain/valueObjects"

type GetShopStatusQuery struct {
	ID shared_valueobject.ID
}

func (_this *GetShopStatusQuery) Name() string {
	return "GetShopSquery"
}
