package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type TableRepositoryImpl struct {
}

func NewTableRepository() TableRepository {
	return &TableRepositoryImpl{}
}

func (repository *TableRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, table domain.Table) domain.Table {
	SQL := "INSERT INTO tables (number, qr_url) VALUES ($1, $2) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, table.Number, table.QRURL).Scan(&id)
	helper.PanicIfError(err)

	table.ID = id
	return table
}

func (repository *TableRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, table domain.Table) domain.Table {
	SQL := "UPDATE tables SET number = $1, qr_url = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, SQL, table.Number, table.QRURL, table.ID)
	helper.PanicIfError(err)

	return table
}

func (repository *TableRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, table domain.Table) {
	SQL := "DELETE FROM tables WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, table.ID)
	helper.PanicIfError(err)
}

func (repository *TableRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Table, error) {
	SQL := "SELECT id, number, qr_url FROM tables WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	table := domain.Table{}
	if rows.Next() {
		err := rows.Scan(&table.ID, &table.Number, &table.QRURL)
		helper.PanicIfError(err)
		return table, nil
	} else {
		return table, errors.New("table is not found")
	}
}

func (repository *TableRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Table {
	SQL := "SELECT id, number, qr_url FROM tables"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var tables []domain.Table
	for rows.Next() {
		table := domain.Table{}
		err := rows.Scan(&table.ID, &table.Number, &table.QRURL)
		helper.PanicIfError(err)
		tables = append(tables, table)
	}
	return tables
}
