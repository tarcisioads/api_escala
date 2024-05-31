package initializers

import (
  "github.com/tarcisioads/api_escala/config"
  "gorm.io/gorm"
)

var (
  DB *gorm.DB
)

func InitializeDB() {
  DB = config.GetSQLite()
}


