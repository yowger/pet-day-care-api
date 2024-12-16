package handlers

import db "github.com/yowger/pet-day-care-api/internal/db/sqlc"

type UserHandler struct {
	queries *db.Queries
}

func NewUserHandler(queries *db.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}
