package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateOrderItem(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteOrderItem(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
