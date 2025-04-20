package controller

import (
	"net/http"
	"strconv"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/service"
	"github.com/julienschmidt/httprouter"
)

type MenuCategoryControllerImpl struct {
	MenuCategoryService service.MenuCategoryService
}

func NewMenuCategoryController(menuCategoryService service.MenuCategoryService) MenuCategoryController {
	return &MenuCategoryControllerImpl{
		MenuCategoryService: menuCategoryService,
	}
}

func (controller *MenuCategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryCreateRequest := web.CreateMenuCategoryRequest{}
	helper.ReadFromRequestBody(r, &menuCategoryCreateRequest)

	menuCategoryResponse := controller.MenuCategoryService.Create(r.Context(), menuCategoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuCategoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuCategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryUpdateRequest := web.UpdateMenuCategoryRequest{}
	helper.ReadFromRequestBody(r, &menuCategoryUpdateRequest)

	menuCategoryId := p.ByName("menuCategoryId")
	id, err := strconv.Atoi(menuCategoryId)
	helper.PanicIfError(err)

	menuCategoryUpdateRequest.ID = id

	menuCategoryResponse := controller.MenuCategoryService.Update(r.Context(), menuCategoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuCategoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuCategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryId := p.ByName("menuCategoryId")
	id, err := strconv.Atoi(menuCategoryId)
	helper.PanicIfError(err)

	controller.MenuCategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuCategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryId := p.ByName("menuCategoryId")
	id, err := strconv.Atoi(menuCategoryId)
	helper.PanicIfError(err)

	menuCategoryResponse := controller.MenuCategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuCategoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MenuCategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	menuCategoryResponse := controller.MenuCategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuCategoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
