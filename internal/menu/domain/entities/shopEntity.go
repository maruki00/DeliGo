package entities

import "time"

type ShopEntity interface {
	GetId() int
	GetName() string
	GetOpenAt() time.Time
	GetClsoeAt() time.Time
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
