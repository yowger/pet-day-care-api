package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
	"github.com/yowger/pet-day-care-api/pkg/validation"
)

type PetHandler struct {
	queries *db.Queries
}

func NewPetHandler(queries *db.Queries) *PetHandler {
	return &PetHandler{queries: queries}
}

func (petHandler *PetHandler) CreatePetHandler(c echo.Context) error {
	var req db.CreatePetParams

	if err := c.Bind(&req); err != nil {
		log.Println("invalid request: ", err)

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validation.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed."})
	}

	pet, err := petHandler.queries.CreatePet(context.Background(), req)
	if err != nil {
		log.Println("\nfailed to create pet: ", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create pet"})
	}

	return c.JSON(http.StatusCreated, pet)
}

func (petHandler *PetHandler) GetPetsPaginatedHandler(c echo.Context) error {
	pets, err := petHandler.queries.GetPetsPaginated(context.Background(), db.GetPetsPaginatedParams{
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch pets"})
	}

	return c.JSON(http.StatusOK, pets)
}

func (petHandler *PetHandler) GetPetByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	id32 := int32(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid pet ID"})
	}

	pet, err := petHandler.queries.GetPetByID(context.Background(), id32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch pet"})
	}

	return c.JSON(http.StatusOK, pet)
}

// func (petHandler *PetHandler) UpdatePetByIdHandler(c echo.Context) error {
// }

// func (petHandler *PetHandler) DeletePetByIdHandler(c echo.Context) error {
// }
