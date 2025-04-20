package repository

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/model/domain"
)

type MenuCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory) domain.MenuCategory
	Update(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory) domain.MenuCategory
	Delete(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.MenuCategory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuCategory
}
