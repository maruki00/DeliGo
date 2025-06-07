package entity

import (
	"time"
)

type ShopEntity interface {
	GetName() string
	GetStatus() bool
	GetOpenAt() time.Time
	GetCloseAt() time.Time
}
