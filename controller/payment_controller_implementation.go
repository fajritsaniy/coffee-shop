package controller

import (
	"net/http"
	"strconv"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/service"
	"github.com/julienschmidt/httprouter"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paymentCreateRequest := web.CreatePaymentRequest{}
	helper.ReadFromRequestBody(r, &paymentCreateRequest)

	paymentResponse := controller.PaymentService.Create(r.Context(), paymentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PaymentControllerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paymentStatusUpdateRequest := web.UpdatePaymentStatusRequest{}
	helper.ReadFromRequestBody(r, &paymentStatusUpdateRequest)

	orderId := p.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	paymentStatusUpdateRequest.OrderID = id

	paymentResponse := controller.PaymentService.UpdateStatus(r.Context(), paymentStatusUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PaymentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paymentId := p.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	paymentResponse := controller.PaymentService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PaymentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paymentResponse := controller.PaymentService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
