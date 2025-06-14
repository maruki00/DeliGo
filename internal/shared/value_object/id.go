package sharedvo

import "github.com/google/uuid"

type ID string

func Parse(id string) ID {
	_id, err := uuid.Parse(id)
	if err != nil {
		return ID("")
	}
	return ID(_id.String())
}
func NewID() ID {
	id := uuid.New()
	return ID(id.String())
}

func (_this *ID) String() string {
	return string(*_this)
}
