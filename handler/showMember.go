package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ShowMemberHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	member := schemas.Member{}
	if err := db.First(&member, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "member not found")
		return
	}

	sendSuccess(ctx, "show-member", member)
}
