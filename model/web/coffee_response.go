package web

import "github.com/fajri/coffee-api/model/domain"

type TableResponse struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	QRURL  string `json:"qr_url"`
}

type MenuCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MenuItemResponse struct {
	ID          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	IsAvailable bool    `json:"is_available"`
	ImageURL    string  `json:"image_url"`
}

type OrderDetailResponse struct {
	ID            int                `json:"id"`
	TableID       int                `json:"table_id"`
	Status        string             `json:"status"`
	PaymentStatus string             `json:"payment_status"`
	Total         float64            `json:"total"`
	OrderItems    []domain.OrderItem `json:"order_items"`
	CreatedAt     string             `json:"created_at"`
}

type OrderResponse struct {
	ID            int     `json:"id"`
	TableID       int     `json:"table_id"`
	Status        string  `json:"status"`
	PaymentStatus string  `json:"payment_status"`
	Total         float64 `json:"total"`
	CreatedAt     string  `json:"created_at"`
}

type OrderItemResponse struct {
	ID       int     `json:"id"`
	OrderID  int     `json:"order_id"`
	MenuID   int     `json:"menu_id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type PaymentResponse struct {
	ID               int     `json:"id"`
	OrderID          int     `json:"order_id"`
	PaymentGateway   string  `json:"payment_gateway"`
	PaymentReference string  `json:"payment_reference"`
	PaidAt           string  `json:"paid_at"`
	Amount           float64 `json:"amount"`
	Status           string  `json:"status"`
}
