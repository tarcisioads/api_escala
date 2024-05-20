package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ShowMinistryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	ministry := schemas.Ministry{}
	if err := db.First(&ministry, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "ministry not found")
		return
	}

	sendSuccess(ctx, "show-ministry", ministry)
}
