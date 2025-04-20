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

type TableServiceImpl struct {
	TableRepository repository.TableRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewTableService(tableRepository repository.TableRepository, DB *sql.DB, validate *validator.Validate) TableService {
	return &TableServiceImpl{
		TableRepository: tableRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *TableServiceImpl) Create(ctx context.Context, request web.CreateTableRequest) web.TableResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	table := domain.Table{
		Number: request.Number,
		QRURL:  request.QRURL,
	}

	table = service.TableRepository.Save(ctx, tx, table)

	return helper.ToTableResponse(table)
}

func (service *TableServiceImpl) Update(ctx context.Context, request web.UpdateTableRequest) web.TableResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	table, err := service.TableRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	table.Number = request.Number
	table.QRURL = request.QRURL

	table = service.TableRepository.Update(ctx, tx, table)

	return helper.ToTableResponse(table)
}

func (service *TableServiceImpl) Delete(ctx context.Context, tableId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	table, err := service.TableRepository.FindById(ctx, tx, tableId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TableRepository.Delete(ctx, tx, table)
}

func (service *TableServiceImpl) FindById(ctx context.Context, tableId int) web.TableResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	table, err := service.TableRepository.FindById(ctx, tx, tableId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTableResponse(table)
}

func (service *TableServiceImpl) FindAll(ctx context.Context) []web.TableResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tables := service.TableRepository.FindAll(ctx, tx)

	return helper.ToTableResponses(tables)
}
