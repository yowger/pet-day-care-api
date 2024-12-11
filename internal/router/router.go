package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
)

type AppRouter struct {
	Echo    *echo.Echo
	Queries *db.Queries
}

func SetUpRouter(e *echo.Echo, queries *db.Queries) {
	ar := &AppRouter{Echo: e, Queries: queries}

	ar.setPetController()
}

func (ar *AppRouter) setPetController() {
	ar.Echo.GET("/pets", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"Tom", "Jerry", "Bob"})
	})
}

// func NewAppRouter(queries *db.Queries) *AppRouter {
// 	return &AppRouter{Queries: queries}
//db

// func (ar *AppRouter) Init(e *echo.Echo) {
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.CORS())
// 	e.Use(middleware.RequestID())

// 	handler.RegisterUserRouters()

// 	// ar.registerPetRoutes(e.Group("/pets"))
// 	// ar.registerBreedRoutes(e.Group("/breed"))
// 	// ar.registerUserRoutes(e.Group("/users"))
// }

// func (ar *AppRouter) registerPetRoutes(g *echo.Group) {
// 	g.GET("", func(c echo.Context) error {
// 		petsPaginated, error := ar.Queries.GetPetsPaginated(context.Background(), db.GetPetsPaginatedParams{Limit: 10, Offset: 0})

// 		if error != nil {
// 			log.Println(error)

// 			return c.JSON(http.StatusInternalServerError, "Error getting pets")
// 		}

// 		return c.JSON(http.StatusOK, petsPaginated)
// 	})
// }

// func (ar *AppRouter) registerUserRoutes(g *echo.Group) {
// 	g.GET("", GetUserHandler)
// }
// func (ar *AppRouter) registerBreedRoutes(g *echo.Group) {
// 	g.GET("", func(c echo.Context) error {
// 		breed, err := ar.Queries.GetAllBreedsPaginated(context.Background(), db.GetAllBreedsPaginatedParams{Limit: 10, Offset: 0})

// 		if err != nil {
// 			log.Println(err)
// 		}

// 		return c.JSON(http.StatusOK, breed)
// 	})
// }

// func GetPetsHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, []string{"Tom", "Jerry", "Bob"})
// }

// func GetUserHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, []string{"Tom", "Jerry", "Bob"})
// }
