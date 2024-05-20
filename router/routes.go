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

    v1.GET("/ministry/:id", handler.ShowMinistryHandler)
    v1.POST("/ministry", handler.CreateMinistryHandler)
    v1.PUT("/ministry/:id", handler.UpdateMinistryHandler)
    v1.DELETE("/ministry/:id", handler.DeleteMinistryHandler)
    v1.GET("/ministries", handler.ListMinistriesHandler)

    v1.GET("/member/:id", handler.ShowMemberHandler)
    v1.POST("/member", handler.CreateMemberHandler)
    v1.PUT("/member/:id", handler.UpdateMemberHandler)
    v1.DELETE("/member/:id", handler.DeleteMemberHandler)
    v1.GET("/members", handler.ListMembersHandler)
 
  }

}
