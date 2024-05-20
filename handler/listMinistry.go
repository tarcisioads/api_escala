package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ListMinistriesHandler(ctx *gin.Context) {
	ministries := []schemas.Ministry{}

	if err := db.Find(&ministries).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing ministries")
		return
	}

	sendSuccess(ctx, "list-ministries", ministries)
}