package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/fajri/coffee-api/exception"
	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/repository"
	"github.com/go-playground/validator"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, DB *sql.DB, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *PaymentServiceImpl) Create(ctx context.Context, request web.CreatePaymentRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Save Payment
	payment := domain.Payment{
		OrderID:          request.OrderID,
		PaymentGateway:   request.PaymentGateway,
		PaymentReference: request.PaymentReference,
		Amount:           request.Amount,
		Status:           "Pending",
		PaidAt:           time.Now(),
	}
	payment = service.PaymentRepository.Save(ctx, tx, payment)
	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) UpdateStatus(ctx context.Context, request web.UpdatePaymentStatusRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment, err := service.PaymentRepository.FindById(ctx, tx, request.OrderID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	payment.Status = request.Status

	payment = service.PaymentRepository.UpdateStatus(ctx, tx, payment)

	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) FindById(ctx context.Context, id int) web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment, err := service.PaymentRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) FindAll(ctx context.Context) []web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payments := service.PaymentRepository.FindAll(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPaymentResponses(payments)
}
