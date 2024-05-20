package handler

import (
	"fmt"
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteMemberHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	member := schemas.Member{}
	if err := db.First(&member, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("member with id: %s not found", id))
		return
	}

	if err := db.Delete(&member).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting member with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-member", member)
}
