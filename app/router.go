package app

import (
	"github.com/fajri/coffee-api/controller"
	"github.com/fajri/coffee-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(tableController controller.TableController) *httprouter.Router {
	router := httprouter.New()

	// Table
	router.GET("/api/v1/tables", tableController.FindAll)
	router.GET("/api/v1/tables/:tableId", tableController.FindById)
	router.POST("/api/v1/tables", tableController.Create)
	router.PUT("/api/v1/tables/:tableId", tableController.Update)
	router.DELETE("/api/v1/tables/:tableId", tableController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
