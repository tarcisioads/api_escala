package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ShowMinistryHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	ministry := schemas.Ministry{}

	if err := db.Model(&schemas.Ministry{}).Preload("Members").First(&ministry, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "ministry not found")
		return
	}

	SendSuccess(ctx, "show-ministry", ministry)
}
