package controller

import (
	"net/http"
	"strconv"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/service"
	"github.com/julienschmidt/httprouter"
)

type MenuItemControllerImpl struct {
	MenuItemService service.MenuItemService
}

func NewMenuItemController(menuItemService service.MenuItemService) MenuItemController {
	return &MenuItemControllerImpl{
		MenuItemService: menuItemService,
	}
}

func (controller *MenuItemControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuItemCreateRequest := web.CreateMenuItemRequest{}
	helper.ReadFromRequestBody(r, &menuItemCreateRequest)

	menuItemResponse := controller.MenuItemService.Create(r.Context(), menuItemCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuItemControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuItemUpdateRequest := web.UpdateMenuItemRequest{}
	helper.ReadFromRequestBody(r, &menuItemUpdateRequest)

	menuItemId := p.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	menuItemUpdateRequest.ID = id

	menuItemResponse := controller.MenuItemService.Update(r.Context(), menuItemUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuItemControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuItemId := p.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	controller.MenuItemService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuItemControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuItemId := p.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	menuItemResponse := controller.MenuItemService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuItemControllerImpl) FindByCategoryID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryId := p.ByName("menuCategoryId")
	id, err := strconv.Atoi(menuCategoryId)
	helper.PanicIfError(err)

	menuItemResponse := controller.MenuItemService.FindByCategoryID(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuItemControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuItemResponse := controller.MenuItemService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
