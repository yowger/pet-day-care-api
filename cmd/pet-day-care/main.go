package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	config "github.com/yowger/pet-day-care-api/config"
	"github.com/yowger/pet-day-care-api/internal/router"
	database "github.com/yowger/pet-day-care-api/pkg/db"
)

func main() {
	cfgPath := "."
	cfgFile, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	pgxPool, err := database.NewPGXPool(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pgxPool.Close()

	e := echo.New()

	router.Init(e)
	// queries.GetAllBreedsPaginated(context.Background(), db.GetAllBreedsPaginatedParams{Limit: 10, Offset: 0})

	log.Println("Server started...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Println("Server started on port 8080...")
	<-ctx.Done()

	log.Println("Shutting down gracefully...")
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}
}
