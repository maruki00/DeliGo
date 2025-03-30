package repositories

import (
	"sync"
)

type ProfileRepository struct {
	sync.Mutex
}
