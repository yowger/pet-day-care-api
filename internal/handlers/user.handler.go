package handlers

import (
	"context"
	"log"
	"net/http"

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

	if c.Bind(&req) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validation.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
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
		log.Println("\nfailed to create user: ", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	formattedCreatedAt := user.CreatedAt.Time.Local().Format("2006-01-02 15:04:05")
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
