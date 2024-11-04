package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once   sync.Once
)

func InitDB() *pgxpool.Pool {
	once.Do(func() {
		dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

		config, err := pgxpool.ParseConfig(dbURL)
		if err != nil {
			// DB_URL not found
			log.Fatal("DB_URL environment variable not set")
		}

		// Set QueryExecMode to avoid automatic statement caching
		config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

		dbPool, err = pgxpool.NewWithConfig(context.Background(), config)

		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})

	return dbPool
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}
