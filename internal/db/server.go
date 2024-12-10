package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api/config"
	"github.com/yowger/pet-day-care-api/internal/router"
	database "github.com/yowger/pet-day-care-api/pkg/db"
)

type Server struct {
	Config   *config.Config
	Echo     *echo.Echo
	PGXPool  *pgxpool.Pool
	WG       *sync.WaitGroup
	Shutdown context.CancelFunc
}

func NewServer(config *config.Config, shutdown context.CancelFunc) *Server {
	pgxPool := setupPGXPool(config)
	server := initEchoServer()
	wg := &sync.WaitGroup{}

	return &Server{
		Config:   config,
		Echo:     server,
		PGXPool:  pgxPool,
		WG:       wg,
		Shutdown: shutdown,
	}
}

func setupPGXPool(config *config.Config) *pgxpool.Pool {
	pgxPool, err := database.NewPGXPool(config.DATABASE_URL)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return pgxPool
}

func initEchoServer() *echo.Echo {
	e := echo.New()

	router.Init(e)

	return e
}

func (server *Server) StartServer() {
	server.WG.Add(1)

	go func() {
		defer server.WG.Done()

		port := fmt.Sprintf(":%s", server.Config.PORT)
		if err := server.Echo.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
}

func (server *Server) HealthCheck(Interval time.Duration, ctx context.Context) {
	server.WG.Add(1)

	go func() {
		defer server.WG.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(Interval):
				if err := server.PGXPool.Ping(context.Background()); err != nil {
					log.Fatalf("Error connecting to database: %v", err)
				}
			}
		}
	}()
}

func (server *Server) GracefulShutdown() {
	log.Println("Shutting down gracefully...")

	server.Shutdown()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Echo.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	server.WG.Wait()

	log.Println("Shutdown complete...")
}
