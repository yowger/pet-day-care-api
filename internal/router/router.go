package router

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

type AppRouter struct {
	Queries *db.Queries
}

func NewAppRouter(queries *db.Queries) *AppRouter {
	return &AppRouter{Queries: queries}
}

func (ar *AppRouter) Init(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	ar.registerPetRoutes(e.Group("/pets"))
	ar.registerUserRoutes(e.Group("/users"))
}

func (ar *AppRouter) registerPetRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		petsPaginated, error := ar.Queries.GetPetsWithOwnersPaginated(context.Background(), db.GetPetsWithOwnersPaginatedParams{Limit: 10, Offset: 0})

		if error != nil {
			log.Println(error)

			return c.JSON(http.StatusInternalServerError, "Error getting pets")
		}

		return c.JSON(http.StatusOK, petsPaginated)
	})
}

func (ar *AppRouter) registerUserRoutes(g *echo.Group) {
	g.GET("", GetUserHandler)
}

func GetPetsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Tom", "Jerry", "Bob"})
}

func GetUserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Tom", "Jerry", "Bob"})
}
