package pkgTenant

import (
	pkgPostgres "deligo/pkg/postgres"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

const CACHE_LIFE = time.Minute * 10

type Tenant[T pkgPostgres.PGDB] struct {
	sync.Mutex
	Pgs     map[string]T
	started time.Time
}

func NewTanenet[T pkgPostgres.PGDB]() *Tenant[T] {
	return &Tenant[T]{
		Pgs:     make(map[string]T, 0),
		started: time.Now(),
	}
}

func (t *Tenant[T]) isExpired() bool {
	return time.Now().After(t.started.Add(CACHE_LIFE))
}

func (t *Tenant[T]) Register(tenant string, userID string) (any, error) {

	t.Lock()
	defer t.Unlock()
	if t.isExpired() {
		t.Pgs = make(map[string]T, 0)
		t.started = time.Now()
	}

	if obj, ok := t.Pgs[tenant]; ok {
		return obj, nil
	}

	url := strings.Trim(fmt.Sprintf("%s/%s", os.Getenv("TANENT_API"), ""), "/")
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err

	}
	data = data["data"].(map[string]any)
	fmt.Println(data, tenant)

	dsn, ok := data[tenant]
	if !ok {
		return nil, errors.New("unauthorized for this tenant or not tenant not found")
	}
	pg, err := pkgPostgres.NewDB[*gorm.DB](dsn.(string))
	if err != nil {
		return nil, err

	}
	t.Pgs[tenant] = any(pg).(T)
	return pg, nil
}

func (t *Tenant[T]) Clean() {
	slog.Info("\ncleaning db objects ....")
	// for _, pg := range t.Pgs {
	// 	_ = pg.Close()
	// }
}
