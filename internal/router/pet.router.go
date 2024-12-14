package router

import (
	"github.com/yowger/pet-day-care-api/internal/handlers"
)

func SetupPetRoutes(appRouter *AppRouter) {
	petHandler := handlers.NewPetHandler(appRouter.Queries)

	publicRoutes := appRouter.Echo.Group("/pets")

	publicRoutes.POST("", petHandler.CreatePetHandler)
	publicRoutes.GET("", petHandler.GetPetsPaginatedHandler)
	publicRoutes.GET("/:id", petHandler.GetPetByIdHandler)
}
