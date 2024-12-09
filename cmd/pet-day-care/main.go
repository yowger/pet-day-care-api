package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/yowger/pet-day-care-api/config"
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

	database, err := database.NewPGXPool(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Close()

	log.Println("Server started...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")
}
