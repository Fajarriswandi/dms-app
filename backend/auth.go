package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// In-memory user store (replace with database in production)
var users = make(map[string]*User)
var usersByUsername = make(map[string]*User)

// Register handles user registration
// @Summary      Register new user
// @Description  Register a new user account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user  body      RegisterRequest  true  "User registration data"
// @Success      201   {object}  AuthResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      409   {object}  ErrorResponse
// @Router       /api/v1/auth/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
		})
		return
	}

	// Validate input
	if req.Username == "" || req.Email == "" || req.Password == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{
			Error:   "validation_error",
			Message: "Username, email, and password are required",
		})
		return
	}

	// Check if user already exists
	if _, exists := usersByUsername[req.Username]; exists {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, ErrorResponse{
			Error:   "user_exists",
			Message: "Username already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to hash password",
		})
		return
	}

	// Create user
	now := time.Now()
	user := &User{
		ID:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Store user
	users[user.ID] = user
	usersByUsername[user.Username] = user

	// Generate JWT token
	token, err := GenerateJWT(user.ID, user.Username)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate token",
		})
		return
	}

	// Return response
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, AuthResponse{
		Token: token,
		User: User{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// Login handles user login
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Login credentials"
// @Success      200          {object}  AuthResponse
// @Failure      401          {object}  ErrorResponse
// @Router       /api/v1/auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request body",
		})
		return
	}

	// Validate input
	if req.Username == "" || req.Password == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{
			Error:   "validation_error",
			Message: "Username and password are required",
		})
		return
	}

	// Find user
	user, exists := usersByUsername[req.Username]
	if !exists {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid username or password",
		})
		return
	}

	// Check password
	if !CheckPasswordHash(req.Password, user.Password) {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid username or password",
		})
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(user.ID, user.Username)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate token",
		})
		return
	}

	// Return response
	render.Status(r, http.StatusOK)
	render.JSON(w, r, AuthResponse{
		Token: token,
		User: User{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// GetProfile returns current user profile
// @Summary      Get user profile
// @Description  Get current authenticated user profile
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  User
// @Failure      401  {object}  ErrorResponse
// @Router       /api/v1/auth/profile [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by JWT middleware)
	userID := r.Context().Value("userID").(string)
	_ = r.Context().Value("username").(string) // Available but not needed for lookup

	// Find user
	user, exists := users[userID]
	if !exists {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found",
		})
		return
	}

	// Return user (without password)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

