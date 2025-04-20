package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type MenuCategoryRepositoryImpl struct {
}

func NewMenuCategoryRepository() MenuCategoryRepository {
	return &MenuCategoryRepositoryImpl{}
}

func (repository *MenuCategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory) domain.MenuCategory {
	SQL := "INSERT INTO menu_category (name) values ($1) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, menuCategory.Name).Scan(&id)
	helper.PanicIfError(err)

	menuCategory.ID = id
	return menuCategory
}

func (repository *MenuCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory) domain.MenuCategory {
	SQL := "UPDATE menu_category SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, menuCategory.Name, menuCategory.ID)
	helper.PanicIfError(err)

	return menuCategory
}

func (repository *MenuCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, menuCategory domain.MenuCategory) {
	SQL := "DELETE FROM menu_category WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, menuCategory.ID)
	helper.PanicIfError(err)
}

func (repository *MenuCategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.MenuCategory, error) {
	SQL := "SELECT id, name FROM menu_category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	menuCategory := domain.MenuCategory{}
	if rows.Next() {
		err := rows.Scan(&menuCategory.ID, &menuCategory.Name)
		helper.PanicIfError(err)
		return menuCategory, nil
	} else {
		return menuCategory, errors.New("menu category is not found")
	}
}

func (repository *MenuCategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuCategory {
	SQL := "SELECT id, name FROM menu_category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var menuCategories []domain.MenuCategory
	for rows.Next() {
		menuCategory := domain.MenuCategory{}
		err := rows.Scan(&menuCategory.ID, &menuCategory.Name)
		helper.PanicIfError(err)
		menuCategories = append(menuCategories, menuCategory)
	}
	return menuCategories
}
