package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/natchapol2347/woodman_web_back/internal/config"
)

func NewPostgres(cfg config.Database) (*sql.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	fmt.Println("Connecting to PostgreSQL...")
	fmt.Println("Connection string:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database because %s", err)
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Database cannot be reached because %s", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	teardown := func() {
		_ = db.Close()
	}

	return db, teardown
}
