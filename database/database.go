package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetDB(host string, port int, schema string, user string, password string) (*sqlx.DB, error) {
	connection := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, schema)
	db, err := sqlx.Open("mysql", connection)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %v", err)
	}

	return db, nil
}

func Migrate(db *sqlx.DB) error {
	migrations := getMigrations()
	for _, migration := range migrations[MigrationIndex:] {
		if err := migration(db); err != nil {
			return fmt.Errorf("migration failed: %v", err)
		}
	}

	return nil
}

func Seed(db *sqlx.DB) error {
	if Seeded {
		return nil
	}

	seeders := getSeeders()
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return fmt.Errorf("seeder failed: %v", err)
		}
	}

	return nil
}
