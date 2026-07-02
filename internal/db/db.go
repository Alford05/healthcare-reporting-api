package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			return db, nil
		}

		fmt.Println("waiting for database...retrying")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("database never became ready: %w", err)
}
