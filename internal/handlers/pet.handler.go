package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

type PetHandler struct {
	queries *db.Queries
}

func NewPetHandler(queries *db.Queries) *PetHandler {
	return &PetHandler{queries: queries}
}

func (h *PetHandler) GetPetsHandler(c echo.Context) error {
	pets, err := h.queries.GetPetsPaginated(context.Background(), db.GetPetsPaginatedParams{
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch pets"})
	}

	return c.JSON(http.StatusOK, pets)
}
