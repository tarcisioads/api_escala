package config

import (
  "fmt"

  "gorm.io/gorm"
)

var (
  db *gorm.DB
  logger *Logger
  basePath string
)

const (
	TestMode = "test"
	ProductionMode = "production"
)

func Init(mode string) error {
  switch mode {
    case TestMode:
      mode = "test"
    default:
      mode = ProductionMode 
  }

	basePath = "/api/v1"

  var err error
  db, err = InitializeSQLite(mode)

  if err != nil {
    return fmt.Errorf("error initializing sqlite: %v", err)
  }

  return nil
}


func GetSQLite() *gorm.DB {
  return db
}

func GetLogger(name string) *Logger {
  logger = NewLogger(name)
  return logger
}

func GetBasePath() string {
  return basePath
}
