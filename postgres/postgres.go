package postgres

import (
	"context"
	"log"

	"github.com/daniarmas/gographqltwitter/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, config *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(config.Database.URL)
	if err != nil {
		log.Fatalf("can't parse database config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	db := &DB{Pool: pool}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("can't ping database: %v", err)
	}
	log.Println("database connected")
}

func (db *DB) Close(ctx context.Context) {
	db.Pool.Close()
}
