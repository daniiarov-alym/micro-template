package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/daniiarov-alym/micro-template/src/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	logger "github.com/sirupsen/logrus"
)

func StartDatabase(ctx context.Context) *pgxpool.Pool {
	createDatabase(ctx)
	migrateDb(ctx)
	cfg := config.Config()
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	pool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		logger.Fatal("Can not connect to database: " + err.Error())
	}
	return pool
}

func createDatabase(ctx context.Context) {
	cfg := config.Config()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.WithError(err).Fatal("Unable to parse database config")
	}
	repo, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Fatalf("Failed to initialise connection to Postgres Database: %s", err)
	}
	_, err = repo.Exec(ctx, "CREATE DATABASE \""+cfg.DbName+"\"")
	if err != nil && !strings.Contains(err.Error(), "(SQLSTATE 42P04)") {
		logger.Fatal("Failed to create a database: " + err.Error())
	}
}

func migrateDb(ctx context.Context) {
	cfg := config.Config()
	logger.Trace("Starting migrations")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	dbCon, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("failed to connect to db :: ", err)
	}

	driver, err := postgres.WithInstance(dbCon, &postgres.Config{})
	if err != nil {
		logger.Fatal("failed to create postgres abstraction :: ", err)
	}

	srcUrl := "file://migrations"
	migrator, err := migrate.NewWithDatabaseInstance(srcUrl, cfg.DbName, driver)
	if err != nil {
		logger.Fatal("failed to create migrator instance :: ", err)
	}

	err = migrator.Up()
	if err != nil && err.Error() != "no change" {
		logger.Fatal("failed to apply migrations :: ", err)
	}
	logger.Trace("Migration finished successfully")
}
