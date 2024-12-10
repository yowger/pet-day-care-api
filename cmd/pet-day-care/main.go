package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yowger/pet-day-care-api/config"
	"github.com/yowger/pet-day-care-api/internal/db"
)

func main() {

	cfgPath := "."
	cfg := config.InitConfig(cfgPath)

	pgxPool := db.SetupPGXPool(cfg.DATABASE_URL)
	defer pgxPool.Close()

	e := db.InitEchoServer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db.StartServer(e, ":8080")

	<-ctx.Done()

	stop()

	db.GracefulShutdown(e, stop)
}
