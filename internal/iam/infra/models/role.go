package models

import valueobjects "deligo/internal/iam/domain/valueobject"

type Role struct {
	ID   valueobjects.ID
	Name string
}
