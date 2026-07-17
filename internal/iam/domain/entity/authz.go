package entity

import "time"

type AuthzEntity interface {
	GetID() int
	GetName() string
	GetAction() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time

	SetID(ID int)
	SetName(Name string)
	SetAction(Action string)
	SetDescription(Description string)
	SetCreatedAt(CreatedAt time.Time)
	SetUpdatedAt(UpdatedAt time.Time)
}
