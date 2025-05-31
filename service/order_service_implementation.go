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

type OrderServiceImpl struct {
	OrderRepository     repository.OrderRepository
	OrderItemRepository repository.OrderItemRepository
	MenuItemRepository  repository.MenuItemRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, orderItemRepository repository.OrderItemRepository, menuItemRepository repository.MenuItemRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository:     orderRepository,
		OrderItemRepository: orderItemRepository,
		MenuItemRepository:  menuItemRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.CreateOrderRequest) web.OrderDetailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Save Order
	order := domain.Order{
		TableID:       request.TableID,
		Status:        "PENDING",
		PaymentStatus: "UNPAID",
		Total:         0,
		CreatedAt:     time.Now(),
	}

	order = service.OrderRepository.Save(ctx, tx, order)

	// Get Order
	orderItems := []domain.OrderItem{}
	total := 0
	// Save Order Items
	for _, item := range request.Items {
		menuItem, err := service.MenuItemRepository.FindById(ctx, tx, item.MenuID)
		helper.PanicIfError(err)
		orderItem := domain.OrderItem{
			OrderID:  order.ID,
			MenuID:   item.MenuID,
			Price:    menuItem.Price,
			Quantity: item.Quantity,
		}
		total += int(orderItem.Price * float64(item.Quantity))
		orderItem = service.OrderItemRepository.Save(ctx, tx, orderItem)
		orderItems = append(orderItems, orderItem)
	}

	// Update total amount
	order.Total = float64(total)
	order = service.OrderRepository.Update(ctx, tx, order)

	return helper.ToOrderDetailResponse(order, orderItems)
}

func (service *OrderServiceImpl) UpdateOrder(ctx context.Context, request web.UpdateOrderRequest) web.OrderDetailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Get Order
	order, err := service.OrderRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Get Order Items
	orderItems, err := service.OrderItemRepository.FindByOrderId(ctx, tx, order.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	order.TableID = request.TableID
	order.Status = request.Status
	order.PaymentStatus = request.PaymentStatus
	order.Total = request.Total
	service.OrderRepository.Update(ctx, tx, order)

	return helper.ToOrderDetailResponse(order, orderItems)
}

func (service *OrderServiceImpl) UpdateOrderItem(ctx context.Context, request web.UpdateOrderItemRequest) web.OrderDetailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Get Order Items
	orders, err := service.OrderRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Get Order
	orderItemsBefore, err := service.OrderItemRepository.FindByOrderId(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	orderItems := orderItemsBefore
	total := orders.Total
	// Save Order Items
	for _, item := range request.Items {
		menuItem, err := service.MenuItemRepository.FindById(ctx, tx, item.MenuID)
		helper.PanicIfError(err)
		orderItem := domain.OrderItem{
			OrderID:  request.ID,
			MenuID:   item.MenuID,
			Price:    menuItem.Price,
			Quantity: item.Quantity,
		}
		total += float64(menuItem.Price) * float64(item.Quantity)
		orderItem = service.OrderItemRepository.Save(ctx, tx, orderItem)
		orderItems = append(orderItems, orderItem)
	}

	// Update Order Total
	orders.Total = total
	orders = service.OrderRepository.Update(ctx, tx, orders)

	return helper.ToOrderDetailResponse(orders, orderItems)
}

func (service *OrderServiceImpl) Delete(ctx context.Context, orderId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderRepository.Delete(ctx, tx, order)

}

func (service *OrderServiceImpl) DeleteOrderItem(ctx context.Context, orderItemId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItem, err := service.OrderItemRepository.FindById(ctx, tx, orderItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrderItemRepository.Delete(ctx, tx, orderItem)

}

func (service *OrderServiceImpl) FindById(ctx context.Context, orderId int) web.OrderDetailResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Get Order
	order, err := service.OrderRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Get Order Items
	orderItems := []domain.OrderItem{}
	orderItems, err = service.OrderItemRepository.FindByOrderId(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderDetailResponse(order, orderItems)
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrderRepository.FindAll(ctx, tx)

	return helper.ToOrderResponses(orders)
}
