package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ShowEscalaHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	escala := schemas.Escala{}
	if err := db.First(&escala, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "escala not found")
		return
	}

	SendSuccess(ctx, "show-escala", escala)
}
