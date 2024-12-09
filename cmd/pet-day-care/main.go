package main

import (
	"fmt"
	"log"

	config "github.com/yowger/pet-day-care-api/config"
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

	fmt.Printf("config: %s\n", cfg)
}
