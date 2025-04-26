package userQueries

import "github.com/google/uuid"

type FindUserByIdQuery struct {
	Key string
	ID  uuid.UUID
}

func (_this *FindUserByIdQuery) Name() string {
	return "FindUserByIdQuery"
}
