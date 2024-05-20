package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateMinistryHandler(ctx *gin.Context) {
	request := UpdateMinistryRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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
	// Update ministry
	if request.Name != "" {
		ministry.Name = request.Name
	}

	// Save ministry
	if err := db.Save(&ministry).Error; err != nil {
		logger.Errorf("error updating ministry: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating ministry")
		return
	}
	sendSuccess(ctx, "update-ministry", ministry)
}
