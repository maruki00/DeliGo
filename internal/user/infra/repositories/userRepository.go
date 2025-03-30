package repositories

import (
	"sync"
)

type UserRepository struct {
	sync.Mutex
}
