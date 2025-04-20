package app

import (
	"database/sql"
	"time"

	"github.com/fajri/coffee-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://fajri:@localhost:5432/coffee_app_test?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
