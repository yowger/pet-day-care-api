package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

type PetHandler struct {
	queries *db.Queries
}

func NewPetHandler(queries *db.Queries) *PetHandler {
	return &PetHandler{queries: queries}
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
