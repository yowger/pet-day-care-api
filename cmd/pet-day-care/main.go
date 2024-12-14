package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yowger/pet-day-care-api/config"
	server "github.com/yowger/pet-day-care-api/internal/db"
	"github.com/yowger/pet-day-care-api/internal/db/middleware"
	db "github.com/yowger/pet-day-care-api/internal/db/sqlc"
	"github.com/yowger/pet-day-care-api/internal/router"
)

func main() {
	configPath := "."
	config := config.InitConfig(configPath)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := server.NewServer(config, stop)
	server.StartServer()
	server.HealthCheck(30*time.Second, ctx)

	middleware.SetupMiddleware(server.Echo)

	queries := db.New(server.PGXPool)
	router.SetupRouter(server.Echo, queries)

	<-ctx.Done()

	server.GracefulShutdown()
}
