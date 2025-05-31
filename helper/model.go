package helper

import (
	"github.com/fajri/coffee-api/model/domain"
	"github.com/fajri/coffee-api/model/web"
)

// Table
func ToTableResponse(table domain.Table) web.TableResponse {
	return web.TableResponse{
		ID:     table.ID,
		Number: table.Number,
		QRURL:  table.QRURL,
	}
}

func ToTableResponses(tables []domain.Table) []web.TableResponse {
	var tableResponses []web.TableResponse
	for _, table := range tables {
		tableResponses = append(tableResponses, ToTableResponse(table))
	}

	return tableResponses
}

// Menu Category
func ToMenuCategoryResponse(menuCategory domain.MenuCategory) web.MenuCategoryResponse {
	return web.MenuCategoryResponse{
		ID:   menuCategory.ID,
		Name: menuCategory.Name,
	}
}

func ToMenuCategoryResponses(menuCategories []domain.MenuCategory) []web.MenuCategoryResponse {
	var menuCategoryResponses []web.MenuCategoryResponse
	for _, menuCateogory := range menuCategories {
		menuCategoryResponses = append(menuCategoryResponses, ToMenuCategoryResponse(menuCateogory))
	}

	return menuCategoryResponses
}

// Menu Item
func ToMenuItemResponse(menuItem domain.MenuItem) web.MenuItemResponse {
	return web.MenuItemResponse{
		ID:          menuItem.ID,
		CategoryID:  menuItem.CategoryID,
		Name:        menuItem.Name,
		Price:       menuItem.Price,
		Description: menuItem.Description,
		IsAvailable: menuItem.IsAvailable,
		ImageURL:    menuItem.ImageURL,
	}
}

func ToMenuItemResponses(menuItems []domain.MenuItem) []web.MenuItemResponse {
	var menuItemResponses []web.MenuItemResponse
	for _, menuCateogory := range menuItems {
		menuItemResponses = append(menuItemResponses, ToMenuItemResponse(menuCateogory))
	}

	return menuItemResponses
}

// Order Detail
func ToOrderDetailResponse(order domain.Order, orderItems []domain.OrderItem) web.OrderDetailResponse {
	return web.OrderDetailResponse{
		ID:            order.ID,
		TableID:       order.TableID,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		Total:         order.Total,
		OrderItems:    orderItems,
		CreatedAt:     order.CreatedAt.Format("2006-01-02T15:04:05-07:00"),
	}
}

// Order
func ToOrderResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		ID:            order.ID,
		TableID:       order.TableID,
		Status:        order.Status,
		PaymentStatus: order.PaymentStatus,
		Total:         order.Total,
		CreatedAt:     order.CreatedAt.Format("2006-01-02T15:04:05-07:00"),
	}
}

func ToOrderResponses(orders []domain.Order) []web.OrderResponse {
	var orderResponses []web.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}

	return orderResponses
}

// Order Item
func ToOrderItemResponse(orderItem domain.OrderItem) web.OrderItemResponse {
	return web.OrderItemResponse{
		ID:       orderItem.ID,
		OrderID:  orderItem.OrderID,
		MenuID:   orderItem.MenuID,
		Price:    orderItem.Price,
		Quantity: orderItem.Quantity,
	}
}

// Payment
func ToPaymentResponse(payment domain.Payment) web.PaymentResponse {
	return web.PaymentResponse{
		ID:               payment.ID,
		OrderID:          payment.OrderID,
		PaymentGateway:   payment.PaymentGateway,
		PaymentReference: payment.PaymentReference,
		PaidAt:           payment.PaidAt.Format("2006-01-02T15:04:05-07:00"),
		Amount:           payment.Amount,
		Status:           payment.Status,
	}
}

func ToPaymentResponses(payments []domain.Payment) []web.PaymentResponse {
	var paymentResponses []web.PaymentResponse
	for _, payment := range payments {
		paymentResponses = append(paymentResponses, ToPaymentResponse(payment))
	}

	return paymentResponses
}
