package web

type UpdateTableRequest struct {
	ID     int    `validate:"required,max=200,min=1" json:"id"`
	Number int    `validate:"required,max=200,min=1" json:"number"`
	QRURL  string `json:"qr_url"`
}

type UpdateMenuCategoryRequest struct {
	ID   int    `validate:"required,max=200,min=1" json:"id"`
	Name string `json:"name"`
}

type UpdateMenuItemRequest struct {
	ID          int     `validate:"required,max=200,min=1" json:"id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	IsAvailable bool    `json:"is_available"`
	ImageURL    string  `json:"image_url"`
}

type UpdateOrderRequest struct {
	ID            int     `validate:"required,max=200,min=1" json:"id"`
	TableID       int     `json:"table_id"`
	Status        string  `json:"status"`
	PaymentStatus string  `json:"payment_status"`
	Total         float64 `json:"total"`
}

type UpdateOrderItemRequest struct {
	ID    int                    `validate:"required,max=200,min=1" json:"id"`
	Items []CreateOrderItemInput `json:"items"`
}

type UpdatePaymentStatusRequest struct {
	OrderID int    `validate:"required,max=200,min=1" json:"order_id"`
	Status  string `json:"status"`
}
