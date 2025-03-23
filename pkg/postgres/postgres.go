package pkgPostgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

const (
	MAX_TRIES = 3
	TIMEOUT   = 1
)

type PGHandler struct {
	//DB       *pgx.Conn
	DB       *sql.DB
	MaxTries int
	Timeout  int
}

func (pg *PGHandler) SetDB(db *sql.DB) {
	pg.DB = db
}

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func NewPG(dsn string) (*PGHandler, error) {

	pg := &PGHandler{
		MaxTries: MAX_TRIES,
		Timeout:  TIMEOUT,
		DB:       nil,
	}
	if dsn == "" {
		return pg, nil
	}
	db, err := NewDB(dsn)
	if err != nil {
		return nil, err
	}
	pg.DB = db
	slog.Info("postgres is connected !")
	return pg, nil
}

func (pg *PGHandler) GetDB() *sql.DB {
	return pg.DB
}

func (pg *PGHandler) Close() {
	_ = pg.DB.Close()
}
