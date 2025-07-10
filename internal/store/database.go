package store

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // Importing pgx driver for PostgreSQL
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost, port=5432, user=postgres, password=postgres, dbname=postgres, sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	fmt.Println("Database connection established successfully")
	return db, nil
}
