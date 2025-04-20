package web

type CreateTableRequest struct {
	Number int    `validate:"required,max=200,min=1" json:"number"`
	QRURL  string `json:"qr_url"`
}

type CreateMenuCategoryRequest struct {
	Name string `json:"name"`
}

type CreateMenuItemRequest struct {
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	IsAvailable bool    `json:"is_available"`
	ImageURL    string  `json:"image_url"`
}

type CreateOrderRequest struct {
	TableID int                    `json:"table_id"`
	Items   []CreateOrderItemInput `json:"items"`
	Total   float64                `json:"total"`
}

type CreateOrderItemInput struct {
	MenuID   int `json:"menu_id"`
	Quantity int `json:"quantity"`
}

type CreatePaymentRequest struct {
	OrderID          int     `validate:"required,max=200,min=1" json:"order_id"`
	PaymentGateway   string  `json:"payment_gateway"`
	PaymentReference string  `json:"payment_reference"`
	Amount           float64 `json:"amount"`
}
