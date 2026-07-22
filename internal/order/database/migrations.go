package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations connects to the database via a standard connection string
// and runs any pending schema changes found in the migrations target folder.
func RunMigrations(databaseURL string) error {
	log.Println("Database migrations: Checking schema state...")

	// golang-migrate expects a clear connection prefix scheme
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migration driver engine: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Database migrations: Schema up to date. No changes applied.")
			return nil
		}
		return fmt.Errorf("migration execution encountered operational error: %w", err)
	}

	log.Println("Database migrations: Schema updated successfully.")
	return nil
}
