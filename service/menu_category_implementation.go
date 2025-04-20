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

type MenuCategoryServiceImpl struct {
	MenuCategoryRepository repository.MenuCategoryRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewMenuCategoryService(menuCategoryRepository repository.MenuCategoryRepository, DB *sql.DB, validate *validator.Validate) MenuCategoryService {
	return &MenuCategoryServiceImpl{
		MenuCategoryRepository: menuCategoryRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *MenuCategoryServiceImpl) Create(ctx context.Context, request web.CreateMenuCategoryRequest) web.MenuCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuCategory := domain.MenuCategory{
		Name: request.Name,
	}

	menuCategory = service.MenuCategoryRepository.Save(ctx, tx, menuCategory)

	return helper.ToMenuCategoryResponse(menuCategory)
}

func (service *MenuCategoryServiceImpl) Update(ctx context.Context, request web.UpdateMenuCategoryRequest) web.MenuCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuCategory, err := service.MenuCategoryRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	menuCategory.Name = request.Name

	menuCategory = service.MenuCategoryRepository.Update(ctx, tx, menuCategory)

	return helper.ToMenuCategoryResponse(menuCategory)
}

func (service *MenuCategoryServiceImpl) Delete(ctx context.Context, menuCategoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuCategory, err := service.MenuCategoryRepository.FindById(ctx, tx, menuCategoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MenuCategoryRepository.Delete(ctx, tx, menuCategory)
}

func (service *MenuCategoryServiceImpl) FindById(ctx context.Context, menuCategoryId int) web.MenuCategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuCategory, err := service.MenuCategoryRepository.FindById(ctx, tx, menuCategoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMenuCategoryResponse(menuCategory)
}

func (service *MenuCategoryServiceImpl) FindAll(ctx context.Context) []web.MenuCategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menuCategories := service.MenuCategoryRepository.FindAll(ctx, tx)

	return helper.ToMenuCategoryResponses(menuCategories)
}
