package migrator

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migrattionsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "Path to storage")
	flag.StringVar(&migrattionsPath, "migrations-path", "", "Path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Name of migrations table")

	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrattionsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrattionsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no changes to apply")

			return
		}
		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
