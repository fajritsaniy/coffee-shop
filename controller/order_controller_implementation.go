package controller

import (
	"net/http"
	"strconv"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/service"
	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderCreateRequest := web.CreateOrderRequest{}
	helper.ReadFromRequestBody(r, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(r.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) UpdateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderUpdateRequest := web.UpdateOrderRequest{}
	helper.ReadFromRequestBody(r, &orderUpdateRequest)

	orderId := p.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderUpdateRequest.ID = id

	orderResponse := controller.OrderService.UpdateOrder(r.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) UpdateOrderItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderItemUpdateRequest := web.UpdateOrderItemRequest{}
	helper.ReadFromRequestBody(r, &orderItemUpdateRequest)

	orderItemId := p.ByName("orderItemId")
	id, err := strconv.Atoi(orderItemId)
	helper.PanicIfError(err)

	orderItemUpdateRequest.ID = id

	orderItemResponse := controller.OrderService.UpdateOrderItem(r.Context(), orderItemUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderItemResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderId := p.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	controller.OrderService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) DeleteOrderItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderItemId := p.ByName("orderItemId")
	id, err := strconv.Atoi(orderItemId)
	helper.PanicIfError(err)

	controller.OrderService.DeleteOrderItem(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderId := p.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OrderControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderResponses := controller.OrderService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
