package service

import (
	"context"

	"github.com/fajri/coffee-api/model/web"
)

type MenuCategoryService interface {
	Create(ctx context.Context, request web.CreateMenuCategoryRequest) web.MenuCategoryResponse
	Update(ctx context.Context, request web.UpdateMenuCategoryRequest) web.MenuCategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.MenuCategoryResponse
	FindAll(ctx context.Context) []web.MenuCategoryResponse
}
