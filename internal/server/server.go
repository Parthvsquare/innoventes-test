package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	dbSql "innoventes-test/db/sqlc"
	myDb "innoventes-test/internal/database"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   *dbSql.Queries
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	log.Default().Printf("postgresUrl %v", postgresUrl)

	result, err := runDBMigration(os.Getenv("MIGRATION_URL"), postgresUrl)
	if err != nil {
		log.Fatalf("error running migration %v", err)
	}

	log.Default().Printf("migration result %v", result)

	database := myDb.New()
	NewServer := &Server{
		port: port,
		db:   database,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func runDBMigration(migrationURL string, dbSource string) (msg string, err error) {
	migration, err := migrate.New(migrationURL, dbSource)

	if err != nil {
		log.Default().Fatalf("error creating migration %v", err)
		return "", err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Default().Printf("failed to run migrate up %v", err)
		return "", err
	}

	log.Println("migration completed successfully")
	return "migration completed successfully", nil
}
