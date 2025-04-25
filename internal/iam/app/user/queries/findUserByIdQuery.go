package userQueries

import "github.com/google/uuid"

type FindUserByIdQuery struct {
	ID uuid.UUID
}

func (_this *FindUserByIdQuery) Name() string {
	return "FindUserByIdQuery"
}
