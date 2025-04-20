package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	UpdateStatus(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Payment, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Payment
}
