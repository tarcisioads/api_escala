package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ShowMemberHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	member := schemas.Member{}
	if err := db.Model(&schemas.Member{}).Preload("Ministries").First(&member, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "member not found")
		return
	}

	SendSuccess(ctx, "show-member", member)
}
