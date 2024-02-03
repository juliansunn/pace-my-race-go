package main

import (
	"api/api"
	db "api/db/sqlc"
	"api/util"
	"database/sql"
	"fmt"
	"os"

	_ "api/docs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Load config from .env file
	config, err := util.Loadconfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Connect to DB using config values from .env file
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	fmt.Println(dbSource)
	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}

	// run db migrations
	runDBMigration(config.MigrationURL, dbSource)

	// Create new store and server
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Msg("cannot create new migration instance")
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msg("failed to run migration up")
	}

	log.Info().Msg("db migrated successfully")
}
