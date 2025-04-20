package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "INSERT INTO orders (table_id, total) VALUES ($1, $2) RETURNING id,status,payment_status,created_at"
	var id int
	var status string
	var paymentStatus string
	var createdAt time.Time
	err := tx.QueryRowContext(ctx, SQL, order.TableID, order.Total).Scan(&id, &status, &paymentStatus, &createdAt)
	helper.PanicIfError(err)

	order.ID = id
	order.Status = status
	order.PaymentStatus = paymentStatus
	order.CreatedAt = createdAt
	return order
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "UPDATE orders SET table_id = $1, status = $2, payment_status = $3, total = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, order.TableID, order.Status, order.PaymentStatus, order.Total, order.ID)
	helper.PanicIfError(err)

	return order
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, order domain.Order) {
	SQL := "DELETE FROM orders WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, order.ID)
	helper.PanicIfError(err)
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Order, error) {
	SQL := "SELECT id, table_id, status, payment_status, total, created_at FROM orders WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.ID, &order.TableID, &order.Status, &order.PaymentStatus, &order.Total, &order.CreatedAt)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, errors.New("order is not found")
	}
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	SQL := `
		SELECT id, table_id, status, payment_status, total, created_at
		FROM orders
		WHERE created_at >= NOW() - INTERVAL '30 days'
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.ID, &order.TableID, &order.Status, &order.PaymentStatus, &order.Total, &order.CreatedAt)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}
	return orders
}
