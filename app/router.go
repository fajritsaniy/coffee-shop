package app

import (
	"github.com/fajri/coffee-api/controller"
	"github.com/fajri/coffee-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(tableController controller.TableController, menuCategory controller.MenuCategoryController, menuItem controller.MenuItemController) *httprouter.Router {
	router := httprouter.New()

	// Table
	router.GET("/api/v1/tables", tableController.FindAll)
	router.GET("/api/v1/tables/:tableId", tableController.FindById)
	router.POST("/api/v1/tables", tableController.Create)
	router.PUT("/api/v1/tables/:tableId", tableController.Update)
	router.DELETE("/api/v1/tables/:tableId", tableController.Delete)

	// Menu Category
	router.POST("/api/v1/menu-categories", menuCategory.Create)
	router.PUT("/api/v1/menu-categories/:menuCategoryId", menuCategory.Update)
	router.DELETE("/api/v1/menu-categories/:menuCategoryId", menuCategory.Delete)
	router.GET("/api/v1/menu-categories/:menuCategoryId", menuCategory.FindById)
	router.GET("/api/v1/menu-categories", menuCategory.FindAll)

	// Menu Item
	router.POST("/api/v1/menu-items", menuItem.Create)
	router.PUT("/api/v1/menu-items/:menuItemId", menuItem.Update)
	router.DELETE("/api/v1/menu-items/:menuItemId", menuItem.Delete)
	router.GET("/api/v1/menu-items/:menuItemId", menuItem.FindById)
	router.GET("/api/v1/menu-items-by-category/:menuCategoryId", menuItem.FindByCategoryID)
	router.GET("/api/v1/menu-items", menuItem.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
