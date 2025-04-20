package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PaymentController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
