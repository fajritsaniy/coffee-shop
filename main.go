package main

import (
	"fmt"
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

	fmt.Println("Service is running...")
	db := app.NewDB()
	validate := validator.New()
	// Table
	tableRepository := repository.NewTableRepository()
	tableService := service.NewTableService(tableRepository, db, validate)
	tableController := controller.NewTableController(tableService)
	// MenuCategory
	menuCategoryRepository := repository.NewMenuCategoryRepository()
	menuCategoryService := service.NewMenuCategoryService(menuCategoryRepository, db, validate)
	menuCategoryController := controller.NewMenuCategoryController(menuCategoryService)
	// MenuItem
	menuItemRepository := repository.NewMenuItemRepository()
	menuItemService := service.NewMenuItemService(menuItemRepository, db, validate)
	menuItemController := controller.NewMenuItemController(menuItemService)

	router := app.NewRouter(tableController, menuCategoryController, menuItemController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
