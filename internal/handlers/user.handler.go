package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
	"github.com/yowger/pet-day-care-api/pkg/auth"
	"github.com/yowger/pet-day-care-api/pkg/validation"
)

type UserHandler struct {
	queries *db.Queries
}

func NewUserHandler(queries *db.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}

type UserResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	RoleID      int32  `json:"role_id"`
	CreatedAt   string `json:"created_at"`
}

func (userHandler *UserHandler) CreateUserHandler(c echo.Context) error {
	var req db.CreateUserParams

	if err := c.Bind(&req); err != nil {
		log.Println("Invalid request:", err)

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validation.Validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]string, len(validationErrors))
		for i, e := range validationErrors {
			errors[i] = fmt.Sprintf("%s is invalid: %s", e.Field(), e.Tag())
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "message": "Validation failed", "errors": errors})
	}

	existingUser, err := userHandler.queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		log.Println("Failed to fetch user:", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
	}

	if existingUser != (db.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User with this email already exists."})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Println("Failed to hash password:", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user, err := userHandler.queries.CreateUser(context.Background(), db.CreateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
		RoleID:      req.RoleID,
	})
	if err != nil {
		log.Println("failed to create user: ", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	formattedCreatedAt := user.CreatedAt.Time.Local().Format(time.RFC3339)
	userResponse := UserResponse{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		RoleID:      user.RoleID,
		CreatedAt:   formattedCreatedAt,
	}

	return c.JSON(http.StatusCreated, userResponse)
}

// func (userHandler *UserHandler) GetUsersHandler(c echo.Context) error {
// }

// func (userHandler *UserHandler) UpdateUserHandler(c echo.Context) error {
// }

// func (userHandler *UserHandler) RemoveUserHandler(c echo.Context) error {
// }
