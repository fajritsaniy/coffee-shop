package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type MenuItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem
	Update(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem
	Delete(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.MenuItem, error)
	FindByCategoryID(ctx context.Context, tx *sql.Tx, categoryId int) (domain.MenuItem, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuItem
}
