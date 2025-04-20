package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
)

type MenuItemRepositoryImpl struct {
}

func NewMenuItemRepository() MenuItemRepository {
	return &MenuItemRepositoryImpl{}
}

func (repository *MenuItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem {
	SQL := "insert into menu_items(category_id, name, price, description, is_available, image_url) values ($1, $2, $3, $4, $5, $6)"
	result, err := tx.ExecContext(ctx, SQL, menuItem.CategoryID, menuItem.Name, menuItem.Price, menuItem.Description, menuItem.IsAvailable, menuItem.ImageURL)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	menuItem.ID = int(id)
	return menuItem
}

func (repository *MenuItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem {
	SQL := "update menu_items set category_id = $1, name = $2, price = $3, description = $4, is_available = $5, image_url = $6 where id = $7"
	_, err := tx.ExecContext(ctx, SQL, menuItem.CategoryID, menuItem.Name, menuItem.Price, menuItem.Description, menuItem.IsAvailable, menuItem.ImageURL, menuItem.ID)
	helper.PanicIfError(err)

	return menuItem
}

func (repository *MenuItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) {
	SQL := "delete from menu_items where id = $1"
	_, err := tx.ExecContext(ctx, SQL, menuItem.ID)
	helper.PanicIfError(err)
}

func (repository *MenuItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.MenuItem, error) {
	SQL := "select id, category_id, name, price, description, is_available, image_url from menu_items where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	menuItem := domain.MenuItem{}
	if rows.Next() {
		err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price, &menuItem.CategoryID)
		helper.PanicIfError(err)
		return menuItem, nil
	} else {
		return menuItem, errors.New("menu item is not found")
	}
}

func (repository *MenuItemRepositoryImpl) FindByCategoryID(ctx context.Context, tx *sql.Tx, categoryId int) (domain.MenuItem, error) {
	SQL := "select id, category_id, name, price, description, is_available, image_url from menu_items where menu_category_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	menuItem := domain.MenuItem{}
	if rows.Next() {
		err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price, &menuItem.CategoryID)
		helper.PanicIfError(err)
		return menuItem, nil
	} else {
		return menuItem, errors.New("menu item is not found")
	}
}

func (repository *MenuItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuItem {
	SQL := "select id, id, category_id, name, price, description, is_available, image_url from menu_items"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var menuItems []domain.MenuItem
	for rows.Next() {
		menuItem := domain.MenuItem{}
		err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price, &menuItem.CategoryID)
		helper.PanicIfError(err)
		menuItems = append(menuItems, menuItem)
	}
	return menuItems
}
