package handler

import (
	"fmt"
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteMinistryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	ministry := schemas.Ministry{}
	if err := db.First(&ministry, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("ministry with id: %s not found", id))
		return
	}

	if err := db.Delete(&ministry).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting ministry with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-ministry", ministry)
}
