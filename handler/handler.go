package handler

import (
  "github.com/tarcisioads/api_escala/config"
  "gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var (
  logger *config.Logger
  db *gorm.DB
)

func InitializeHandler() {
  logger = config.GetLogger("handler")
  db = config.GetSQLite()
}

func SetUpMockRouter() *gin.Engine {
	router := gin.Default()
  gin.SetMode(gin.TestMode) 
	return router
}
