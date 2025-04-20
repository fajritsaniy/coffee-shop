package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type TableRepository interface {
	Save(ctx context.Context, tx *sql.Tx, table domain.Table) domain.Table
	Update(ctx context.Context, tx *sql.Tx, table domain.Table) domain.Table
	Delete(ctx context.Context, tx *sql.Tx, table domain.Table)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Table, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Table
}
