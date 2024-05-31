package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ListMinistriesHandler(ctx *gin.Context) {
	ministries := []schemas.Ministry{}

	if err := db.Model(&schemas.Ministry{}).Preload("Members").Find(&ministries).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing ministries")
		return
	}

	SendSuccess(ctx, "list-ministries", ministries)
}
