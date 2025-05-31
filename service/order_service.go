package service

import (
	"context"

	"github.com/fajri/coffee-api/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.CreateOrderRequest) web.OrderDetailResponse
	UpdateOrder(ctx context.Context, request web.UpdateOrderRequest) web.OrderDetailResponse
	UpdateOrderItem(ctx context.Context, request web.UpdateOrderItemRequest) web.OrderDetailResponse
	Delete(ctx context.Context, orderId int)
	DeleteOrderItem(ctx context.Context, orderItemId int)
	FindById(ctx context.Context, orderId int) web.OrderDetailResponse
	FindAll(ctx context.Context) []web.OrderResponse
}
