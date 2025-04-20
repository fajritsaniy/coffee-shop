package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type OrderItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem) domain.OrderItem
	Update(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem) domain.OrderItem
	Delete(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem)
	FindById(ctx context.Context, tx *sql.Tx, itemId int) (domain.OrderItem, error)
	FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int) ([]domain.OrderItem, error)
}
