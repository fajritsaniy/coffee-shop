package middleware

import (
	"net/http" // Make sure http is imported

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
	// --- START: CRITICAL CHANGE FOR CORS ---
	// Allow OPTIONS requests to pass through without authentication.
	// Browsers send OPTIONS as a preflight for CORS.
	if r.Method == http.MethodOptions {
		middleware.Handler.ServeHTTP(w, r) // Pass control to the next handler (CORS middleware)
		return                             // Stop processing here for OPTIONS requests
	}
	// --- END: CRITICAL CHANGE FOR CORS ---

	// Your existing authentication logic for other HTTP methods
	if "RAHASIA" == r.Header.Get("X-API-Key") {
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
