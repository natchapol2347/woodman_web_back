package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/natchapol2347/woodman_web_back/internal/config"
)

func NewPostgres(cfg config.Database) (*sql.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database because %s", err)
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Database cannot be reached because %s", err)
	}

	teardown := func() {
		_ = db.Close()
	}

	return db, teardown
}
