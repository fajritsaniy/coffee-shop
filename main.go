package main

import (
	"fmt"
	"log" // <--- Add this import for logging errors properly
	"net/http"

	"github.com/fajri/coffee-api/app"
	"github.com/fajri/coffee-api/controller"
	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/middleware"
	"github.com/fajri/coffee-api/repository"
	"github.com/fajri/coffee-api/service"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"

	"github.com/rs/cors" // <--- Import the cors package
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

	// Order
	orderRepository := repository.NewOrderRepository()
	orderItemRepository := repository.NewOrderItemRepository()
	orderService := service.NewOrderService(orderRepository, orderItemRepository, menuItemRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	// Initialize your main router (from app.NewRouter)
	mainRouter := app.NewRouter(tableController, menuCategoryController, menuItemController, orderController)

	// --- Start CORS Configuration ---
	// Configure the CORS middleware options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                      // Your frontend's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},    // Common HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Api-Key"}, // Common headers, include Authorization if your frontend sends tokens
		AllowCredentials: true,                                                   // Set to true if your frontend sends cookies or authorization headers
		Debug:            true,                                                   // Set to true during development to see CORS logs
	}
	// Create the CORS handler
	corsHandler := cors.New(corsOptions).Handler(mainRouter)
	// --- End CORS Configuration ---

	// Now, wrap the CORS handler with your authentication middleware
	// The request flow will be: CORS -> Auth -> Router
	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(corsHandler), // <--- Use the corsHandler here
	}

	log.Printf("Server listening on %s", server.Addr) // <--- Better logging for server start
	err := server.ListenAndServe()
	helper.PanicIfError(err) // Consider using log.Fatal for more robust error handling in main
}
