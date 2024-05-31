package main

import (
	"github.com/tarcisioads/api_escala/config"
	"github.com/tarcisioads/api_escala/initializers"
	"github.com/tarcisioads/api_escala/router"
)

var (
	logger *config.Logger
)

func init() {
  logger = config.GetLogger("main")
  initializers.LoadEnvVariables()
  initializers.InitializeDB()
}

func main() {
	// Initialize Configs
	err := config.Init(config.ProductionMode)
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
