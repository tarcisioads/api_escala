package handler

import (
  "github.com/tarcisioads/api_escala/config"
  "gorm.io/gorm"
)

var (
  logger *config.Logger
  db *gorm.DB
)

func InitializeHandler() {
  logger = config.GetLogger("handler")
  db = config.GetSQLite()
}
