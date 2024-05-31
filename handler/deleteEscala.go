package handler

import (
	"fmt"
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteEscalaHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	escala := schemas.Escala{}
	if err := db.First(&escala, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("escala with id: %s not found", id))
		return
	}

	if err := db.Delete(&escala).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting escala with id: %s", id))
		return
	}
	SendSuccess(ctx, "delete-escala", escala)
}
