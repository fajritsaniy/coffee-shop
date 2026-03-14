package middleware

import (
	"net/http" // Make sure http is imported
	"os"

	"github.com/fajri/coffee-api/helper"
	"github.com/fajri/coffee-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. Always allow OPTIONS (CORS preflight) and health check
	if r.Method == http.MethodOptions || r.URL.Path == "/api/v1/health" {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	// 2. Allow Cashier functionality (Public access)
	// - All GET requests (menu browsing, categories)
	// - Creating an order (POST /api/v1/orders)
	if r.Method == http.MethodGet || (r.Method == http.MethodPost && r.URL.Path == "/api/v1/orders") {
		middleware.Handler.ServeHTTP(w, r)
		return
	}

	// 3. Admin functionality (Requires RAHASIA API Key)
	// For POST, PUT, DELETE operations on menu items, categories, etc.
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "RAHASIA"
	}

	if apiKey == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
