package service

import (
	"context"

	"github.com/fajri/coffee-api/model/web"
)

type PaymentService interface {
	Create(ctx context.Context, request web.CreatePaymentRequest) web.PaymentResponse
	UpdateStatus(ctx context.Context, request web.UpdatePaymentStatusRequest) web.PaymentResponse
	FindById(ctx context.Context, paymentId int) web.PaymentResponse
	FindAll(ctx context.Context) []web.PaymentResponse
}
