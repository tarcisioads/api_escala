package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ListMembersHandler(ctx *gin.Context) {
	members := []schemas.Member{}

	if err := db.Model(&schemas.Member{}).Preload("Ministries").Find(&members).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing members")
		return
	}

	sendSuccess(ctx, "list-members", members)
}
