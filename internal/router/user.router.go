package router

import "github.com/yowger/pet-day-care-api/internal/handlers"

func SetupUserRoutes(appRouter *AppRouter) {
	petHandler := handlers.NewUserHandler(appRouter.Queries)

	publicRoutes := appRouter.Echo.Group("/users")

	publicRoutes.POST("", petHandler.CreateUserHandler)
	publicRoutes.POST("/login", petHandler.Login)
}
