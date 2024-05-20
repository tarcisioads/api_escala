package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func CreateMinistryHandler(ctx *gin.Context) {
	request := CreateMinistryRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request body : %v", request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ministry := schemas.Ministry{
    Name: request.Name,
	}

	if err := db.Create(&ministry).Error; err != nil {
		logger.Errorf("error creating escala: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating ministry on database")
		return
	}

	sendSuccess(ctx, "create-ministry", ministry)
}
