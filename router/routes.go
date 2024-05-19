package router

import (
  "github.com/tarcisioads/api_escala/handler"
  "github.com/gin-gonic/gin"
)


func initializeRoutes(router *gin.Engine) {
  handler.InitializeHandler()
  basePath := "/api/v1"
  v1 := router.Group(basePath)
  {
    v1.GET("/escala/:id", handler.ShowEscalaHandler)
    v1.POST("/escala", handler.CreateEscalaHandler)
    v1.PUT("/escala/:id", handler.UpdateEscalaHandler)
    v1.DELETE("/escala/:id", handler.DeleteEscalaHandler)
    v1.GET("/escalas", handler.ListEscalasHandler)
  }

}
