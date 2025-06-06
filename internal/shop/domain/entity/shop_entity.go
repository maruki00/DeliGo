package entity

import (
	"time"
)

type ShopEntity interface {
	GetName() string
	GetOpenAt() time.Time
	GetCloseAt() time.Time
}
