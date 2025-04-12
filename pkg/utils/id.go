package pkgUtils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func Id() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ret, err := ulid.New(ms, entropy)
	if err != nil {
		return ""
	}
	return ret.String()
}
