package sharedvo

import "github.com/google/uuid"

type ID string

func NewID() ID {
	id := uuid.New()
	return ID(id.String())
}

func (_this *ID) String() string {
	return string(*_this)
}
