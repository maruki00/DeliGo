package query

import pkgCqrs "deligo/pkg/cqrs"

type GetManyProductQuery struct {
}

func (_this *GetManyProductQuery) Handler(q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
