package db

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api/internal/router"
	database "github.com/yowger/pet-day-care-api/pkg/db"
)

func SetupPGXPool(connString string) *pgxpool.Pool {
	pgxPool, err := database.NewPGXPool(connString)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return pgxPool
}

func InitEchoServer() *echo.Echo {
	e := echo.New()

	router.Init(e)

	return e
}

func StartServer(e *echo.Echo, port string) {
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", port)
}

func GracefulShutdown(e *echo.Echo, stop context.CancelFunc) {
	log.Println("Shutting down gracefully...")

	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}
}
