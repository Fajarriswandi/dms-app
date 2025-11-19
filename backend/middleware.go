package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// JWTAuthMiddleware validates JWT token and adds user info to context
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, ErrorResponse{
				Error:   "unauthorized",
				Message: "Authorization header is required",
			})
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, ErrorResponse{
				Error:   "unauthorized",
				Message: "Invalid authorization header format. Use: Bearer <token>",
			})
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, ErrorResponse{
				Error:   "unauthorized",
				Message: "Invalid or expired token",
			})
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)

		// Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SecurityHeadersMiddleware adds security headers
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

