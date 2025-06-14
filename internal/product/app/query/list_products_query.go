package query

import pkgCqrs "deligo/pkg/cqrs"

type ListProductsQuery struct {
}

func (_this *ListProductsQuery) Handler(q pkgCqrs.Query) (interface{}, error) {

	return nil, nil
}
