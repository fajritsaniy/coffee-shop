package domain

import "time"

// Table represents a restaurant table in the domain layer
type Table struct {
	ID     int
	Number int
	QRURL  string
}

// MenuCategory represents categories for menu items
type MenuCategory struct {
	ID   int
	Name string
}

// MenuItem represents a menu item
type MenuItem struct {
	ID          int
	CategoryID  int
	Name        string
	Price       float64
	Description string
	IsAvailable bool
	ImageURL    string
}

// Order represents a customer order
type Order struct {
	ID            int
	TableID       int
	Status        string
	PaymentStatus string
	Total         float64
	CreatedAt     time.Time
}

type MenuItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID       int              `json:"id"`
	OrderID  int              `json:"order_id"`
	Menu     MenuItemResponse `json:"menu"`
	Price    float64          `json:"price"`
	Quantity int              `json:"quantity"`
	Notes    string           `json:"notes,omitempty"` // Optional notes for the order item
}

// Payment represents a payment transaction
type Payment struct {
	ID               int
	OrderID          int
	PaymentGateway   string
	PaymentReference string
	PaidAt           time.Time
	Amount           float64
	Status           string
}
