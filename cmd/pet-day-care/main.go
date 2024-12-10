package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yowger/pet-day-care-api/config"
	"github.com/yowger/pet-day-care-api/internal/db"
)

func main() {
	var wg sync.WaitGroup

	configPath := "."
	config := config.InitConfig(configPath)

	pgxPool := db.SetupPGXPool(config.DATABASE_URL)
	defer pgxPool.Close()

	e := db.InitEchoServer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg.Add(1)
	db.StartServer(e, ":8080", &wg)

	wg.Add(1)
	go db.HealthCheck(pgxPool, &wg)

	<-ctx.Done()

	db.GracefulShutdown(e, stop)
}
