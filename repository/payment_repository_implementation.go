package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository *PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "insert into payments(order_id, payment_gateway, payment_reference, paid_at, amount, status) values ($1, $2, $3, $4, $5, $6 )"
	result, err := tx.ExecContext(ctx, SQL, payment.OrderID, payment.PaymentGateway, payment.PaymentReference, payment.PaidAt, payment.Amount, payment.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	payment.ID = int(id)
	return payment
}

func (repository *PaymentRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "update payments set status = $1 where id = $2"
	_, err := tx.ExecContext(ctx, SQL, payment.Status, payment.ID)
	helper.PanicIfError(err)

	return payment
}

func (repository *PaymentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Payment, error) {
	SQL := "select id, order_id, payment_gateway, payment_reference, paid_at, amount, status from payments where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	payment := domain.Payment{}
	if rows.Next() {
		err := rows.Scan(&payment.ID, &payment.OrderID, &payment.PaymentGateway, &payment.PaymentReference, &payment.PaidAt, &payment.Amount, &payment.Status)
		helper.PanicIfError(err)
		return payment, nil
	} else {
		return payment, errors.New("payment is not found")
	}
}

func (repository *PaymentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Payment {
	SQL := "select id, order_id, payment_gateway, payment_reference, paid_at, amount, status from payments WHERE paid_at >= NOW() - INTERVAL '30 days'"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var payments []domain.Payment
	for rows.Next() {
		payment := domain.Payment{}
		err := rows.Scan(&payment.ID, &payment.OrderID, &payment.PaymentGateway, &payment.PaymentReference, &payment.PaidAt, &payment.Amount, &payment.Status)
		helper.PanicIfError(err)
		payments = append(payments, payment)
	}

	return payments
}
