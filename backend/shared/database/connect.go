package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func Connect() (*pgx.Conn, error) {

	// Read DATABASE_URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL not set in .env")
	}
	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	// Test the connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}
	return conn, nil
}
