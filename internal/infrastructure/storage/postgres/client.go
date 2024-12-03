package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sample-api/internal/config"
)

type clientPostgres struct {
}

func Connect(config config.DatabaseConfig) *sql.DB {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
