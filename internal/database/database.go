package database

import (
	"context"
	"fmt"
	db "innoventes-test/db/sqlc"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() *db.Queries {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	dbPool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}

	dbConnection := db.New(dbPool)
	return dbConnection
}

// func (s *service) Health() map[string]string {
// 	_, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// err := s.db.PingContext(ctx)
// if err != nil {
// 	log.Fatalf(fmt.Sprintf("db down: %v", err))
// }

// 	return map[string]string{
// 		"message": "It's healthy",
// 	}
// }
