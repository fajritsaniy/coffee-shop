package main

import (
	"net/http"

	"github.com/fajri/coffee-api/app"
	"github.com/fajri/coffee-api/controller"
	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/middleware"
	"github.com/fajri/coffee-api/repository"
	"github.com/fajri/coffee-api/service"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	tableRepository := repository.NewTableRepository()
	tableService := service.NewTableService(tableRepository, db, validate)
	tableController := controller.NewTableController(tableService)

	router := app.NewRouter(tableController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
