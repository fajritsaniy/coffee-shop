package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type OrderItemRepositoryImpl struct {
}

func NewOrderItemRepository() OrderItemRepository {
	return &OrderItemRepositoryImpl{}
}

// 	Save(ctx context.Context, tx *sql.Tx, orderItems []domain.OrderItem) []domain.OrderItem

func (repository *OrderItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem) domain.OrderItem {
	SQL := "INSERT INTO order_items(order_id, menu_id, price, quantity, notes) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, orderItem.OrderID, orderItem.MenuID, orderItem.Price, orderItem.Quantity, orderItem.Notes).Scan(&id)
	helper.PanicIfError(err)

	orderItem.ID = id
	return orderItem
}

func (repository *OrderItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem) domain.OrderItem {
	SQL := "update order_items set order_id = $1, menu_id = $2, price = $3, quantity = $4 where id = $5"
	_, err := tx.ExecContext(ctx, SQL, orderItem.OrderID, orderItem.MenuID, orderItem.Price, orderItem.Quantity, orderItem.ID)
	helper.PanicIfError(err)

	return orderItem
}

func (repository *OrderItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderItem domain.OrderItem) {
	SQL := "DELETE FROM order_items WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, orderItem.ID)
	helper.PanicIfError(err)
}

func (repository *OrderItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.OrderItem, error) {
	SQL := "select id, order_id, menu_id, price, quantity, notes from order_items where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	orderItem := domain.OrderItem{}
	if rows.Next() {
		err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.MenuID, &orderItem.Price, &orderItem.Quantity, &orderItem.Notes)
		helper.PanicIfError(err)
		return orderItem, nil
	} else {
		return orderItem, errors.New("menu item is not found")
	}
}

func (repository *OrderItemRepositoryImpl) FindByOrderId(ctx context.Context, tx *sql.Tx, orderId int) ([]domain.OrderItem, error) {
	SQL := "SELECT id, order_id, menu_id, price, quantity, notes FROM order_items WHERE order_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderItems := []domain.OrderItem{}
	for rows.Next() {
		orderItem := domain.OrderItem{}
		err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.MenuID, &orderItem.Price, &orderItem.Quantity, &orderItem.Notes)
		helper.PanicIfError(err)
		orderItems = append(orderItems, orderItem)
	}
	if len(orderItems) == 0 {
		return []domain.OrderItem{}, errors.New("order items not found")
	}
	return orderItems, nil
}
