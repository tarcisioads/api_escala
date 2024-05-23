package config

import (
  "os"
  "github.com/tarcisioads/api_escala/schemas"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

const (
  dbDefaultPath = "./db/main.db"
  dbTestPath = "./db/test.db"
)

func InitializeSQLite(mode string) (*gorm.DB, error) {
  logger := GetLogger("sqlite")
  logger.Info("database finding...")
  dbPath := dbDefaultPath
  switch mode {
    case TestMode:
      dbPath = dbTestPath
      _, err := os.Stat(dbPath)
      if !os.IsNotExist(err) {
        err := os.Remove(dbPath)
        if err != nil {
          logger.Errorf("sqlite test error: %v", err)
          return nil, err
        }
      } 
  }

  _, err := os.Stat(dbPath)
  if os.IsNotExist(err) {
    logger.Info("database file not found, creating...")
    err = os.MkdirAll("./db", os.ModePerm)
    if err != nil {
      return nil, err
    }
    file, err := os.Create(dbPath)
    if err != nil {
      return nil, err
    }

    file.Close()
  }

  db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
  if err != nil {
    logger.Errorf("sqlite opening error: %v", err)
    return nil, err
  }
  err = db.AutoMigrate(&schemas.Escala{}, &schemas.Ministry{}, &schemas.Member{})
  if err != nil {
    logger.Errorf("sqlite automigration error: %v", err)
    return nil, err
  }

  return db, nil
}
