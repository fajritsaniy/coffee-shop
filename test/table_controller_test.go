package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fajri/coffee-api/app"
	"github.com/fajri/coffee-api/controller"
	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/middleware"
	"github.com/fajri/coffee-api/repository"
	"github.com/fajri/coffee-api/service"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://fajri@localhost:5432/coffee_app_test?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {

	validate := validator.New()
	tableRepository := repository.NewTableRepository()
	tableService := service.NewTableService(tableRepository, db, validate)
	tableController := controller.NewTableController(tableService)
	router := app.NewRouter(tableController)

	return middleware.NewAuthMiddleware(router)
}

func truncateTable(db *sql.DB) {
	db.Exec("TRUNCATE table")
}

func TestCreateTableSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"number":12, "qr_url":"kedaikopi104.com/12"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/tables", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, 12, responseBody["data"].(map[string]interface{})["number"])
	assert.Equal(t, "kedaikopi104.com/12", responseBody["data"].(map[string]interface{})["qr_url"])

}

// func TestCreateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : ""}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD REQUEST", responseBody["status"])
// }

// func TestUpdateCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
// }

// func TestUpdateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : ""}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD REQUEST", responseBody["status"])
// }

// func TestGetCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
// }

// func TestGetCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT FOUND", responseBody["status"])
// }

// func TestDeleteCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// }

// func TestDeleteCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT FOUND", responseBody["status"])
// }

// func TestListCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Computer",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
// 	request.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])

// 	var categories = responseBody["data"].([]interface{})
// 	categoriesResponse1 := categories[0].(map[string]interface{})
// 	categoriesResponse2 := categories[1].(map[string]interface{})
// 	assert.Equal(t, category1.Id, int(categoriesResponse1["id"].(float64)))
// 	assert.Equal(t, category1.Name, categoriesResponse1["name"])
// 	assert.Equal(t, category2.Id, int(categoriesResponse2["id"].(float64)))
// 	assert.Equal(t, category2.Name, categoriesResponse2["name"])
// }

// func TestUnauthorized(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, domain.Category{
// 		Name: "Laptop",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 401, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 401, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
// }
