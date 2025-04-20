package controller

import (
	"net/http"
	"strconv"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/service"
	"github.com/julienschmidt/httprouter"
)

type TableControllerImpl struct {
	TableService service.TableService
}

func NewTableController(tableService service.TableService) TableController {
	return &TableControllerImpl{
		TableService: tableService,
	}
}

func (controller *TableControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableCreateRequest := web.CreateTableRequest{}
	helper.ReadFromRequestBody(r, &tableCreateRequest)

	tableResponse := controller.TableService.Create(r.Context(), tableCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tableResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TableControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableUpdateRequest := web.UpdateTableRequest{}
	helper.ReadFromRequestBody(r, &tableUpdateRequest)

	tableId := p.ByName("tableId")
	id, err := strconv.Atoi(tableId)
	helper.PanicIfError(err)

	tableUpdateRequest.ID = id

	tableResponse := controller.TableService.Update(r.Context(), tableUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tableResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TableControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableId := p.ByName("tableId")
	id, err := strconv.Atoi(tableId)
	helper.PanicIfError(err)

	controller.TableService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TableControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableId := p.ByName("tableId")
	id, err := strconv.Atoi(tableId)
	helper.PanicIfError(err)

	tableResponse := controller.TableService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tableResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TableControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableResponse := controller.TableService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tableResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
