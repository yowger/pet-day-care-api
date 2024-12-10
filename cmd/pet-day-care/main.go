package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yowger/pet-day-care-api/config"
	"github.com/yowger/pet-day-care-api/internal/db"
)

func main() {
	configPath := "."
	config := config.InitConfig(configPath)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := db.NewServer(config, stop)
	defer server.PGXPool.Close()

	server.StartServer()
	server.HealthCheck(30*time.Second, ctx)

	<-ctx.Done()

	server.GracefulShutdown()
}
