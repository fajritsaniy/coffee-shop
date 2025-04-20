package service

import (
	"context"

	"github.com/fajri/coffee-api/model/web"
)

type TableService interface {
	Create(ctx context.Context, request web.CreateTableRequest) web.TableResponse
	Update(ctx context.Context, request web.UpdateTableRequest) web.TableResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.TableResponse
	FindAll(ctx context.Context) []web.TableResponse
}
