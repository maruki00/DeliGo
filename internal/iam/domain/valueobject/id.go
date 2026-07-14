package valueobject

import "github.com/google/uuid"

type ID string

func NewID() ID {
	id := uuid.New()
	return ID(id.String())
}

func GetID(id string) ID {
	if _, err := uuid.Parse(id); err != nil {
		// check if id is empty then not valid id
		return ID("")
	}
	return ID(id)
}

func (_this *ID) String() string {
	return string(*_this)
}
