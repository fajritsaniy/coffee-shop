package service

import (
	"context"

	"github.com/fajri/coffee-api/model/web"
)

type MenuItemService interface {
	Create(ctx context.Context, request web.CreateMenuItemRequest) web.MenuItemResponse
	Update(ctx context.Context, request web.UpdateMenuItemRequest) web.MenuItemResponse
	Delete(ctx context.Context, itemId int)
	FindById(ctx context.Context, itemId int) web.MenuItemResponse
	FindByCategoryID(ctx context.Context, categoryId int) web.MenuItemResponse
	FindAll(ctx context.Context) []web.MenuItemResponse
}
