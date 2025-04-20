package service

import (
	"context"
	"database/sql"

	"github.com/fajri/coffee-api/exception"
	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/domain"
	"github.com/fajri/coffee-api/model/web"
	"github.com/fajri/coffee-api/repository"
	"github.com/go-playground/validator"
)

type MenuItemServiceImpl struct {
	MenuItemRepository repository.MenuItemRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewMenuItemService(menuItemRepository repository.MenuItemRepository, DB *sql.DB, validate *validator.Validate) MenuItemService {
	return &MenuItemServiceImpl{
		MenuItemRepository: menuItemRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *MenuItemServiceImpl) Create(ctx context.Context, request web.CreateMenuItemRequest) web.MenuItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItem := domain.MenuItem{
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		IsAvailable: request.IsAvailable,
		ImageURL:    request.ImageURL,
	}

	menuItem = service.MenuItemRepository.Save(ctx, tx, menuItem)

	return helper.ToMenuItemResponse(menuItem)
}

func (service *MenuItemServiceImpl) Update(ctx context.Context, request web.UpdateMenuItemRequest) web.MenuItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	menuItem.CategoryID = request.CategoryID
	menuItem.Name = request.Name
	menuItem.Price = request.Price
	menuItem.Description = request.Description
	menuItem.IsAvailable = request.IsAvailable
	menuItem.ImageURL = request.ImageURL

	menuItem = service.MenuItemRepository.Update(ctx, tx, menuItem)

	return helper.ToMenuItemResponse(menuItem)
}

func (service *MenuItemServiceImpl) Delete(ctx context.Context, menuItemId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, menuItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MenuItemRepository.Delete(ctx, tx, menuItem)
}

func (service *MenuItemServiceImpl) FindById(ctx context.Context, menuItemId int) web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, menuItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMenuItemResponse(menuItem)
}

func (service *MenuItemServiceImpl) FindByCategoryID(ctx context.Context, categoryId int) web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItem, err := service.MenuItemRepository.FindByCategoryID(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMenuItemResponse(menuItem)
}

func (service *MenuItemServiceImpl) FindAll(ctx context.Context) []web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuItems := service.MenuItemRepository.FindAll(ctx, tx)

	return helper.ToMenuItemResponses(menuItems)
}
