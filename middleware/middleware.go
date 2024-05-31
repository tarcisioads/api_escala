package middleware

import (
  "github.com/tarcisioads/api_escala/config"
)

var (
  logger *config.Logger
)

func InitializeMiddleware() {
  logger = config.GetLogger("middleware")
}


