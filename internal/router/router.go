package router

import (
	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

type AppRouter struct {
	Echo    *echo.Echo
	Queries *db.Queries
}

func SetupRouter(e *echo.Echo, queries *db.Queries) {
	appRouter := &AppRouter{Echo: e, Queries: queries}

	SetupPetRoutes(appRouter)
	SetupUserRoutes(appRouter)
}
