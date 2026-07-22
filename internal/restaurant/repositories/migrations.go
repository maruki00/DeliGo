package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunDatabaseMigrations executes pending .sql migrations against the target database connection string
func RunDatabaseMigrations(databaseURL string, pathMigration string) {
	// golang-migrate expects a standard 'postgres://' schema format
	m, err := migrate.New(
		pathMigration, //"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Could not initialize migration engine instance: %v", err)
	}

	fmt.Println("Applying pending database migrations...")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Database schema is already up to date.")
			return
		}
		log.Fatalf("Failed to execute database migration lifecycle: %v", err)
	}
	fmt.Println("Migrations executed successfully.")
}
