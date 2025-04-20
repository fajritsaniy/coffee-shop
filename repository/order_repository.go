package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Delete(ctx context.Context, tx *sql.Tx, order domain.Order)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Order, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Order
}
