package app

import (
	"database/sql"
	"os"
	"time"

	"github.com/fajri/coffee-api/helper"
)

func NewDB() *sql.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUrl := os.Getenv("DB_URL")

	if dbDriver == "" {
		dbDriver = "postgres"
	}
	if dbUrl == "" {
		// Fallback for development if .env is not loaded yet or missing
		dbUrl = "postgres://fajri:@localhost:5432/coffee_app_test?sslmode=disable"
	}

	db, err := sql.Open(dbDriver, dbUrl)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
