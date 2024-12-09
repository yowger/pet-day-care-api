package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

func NewQueries(pool *pgxpool.Pool) *db.Queries {
	return db.New(pool)
}
