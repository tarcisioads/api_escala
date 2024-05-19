package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func CreateEscalaHandler(ctx *gin.Context) {
	request := CreateEscalaRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request body : %v", request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	escala := schemas.Escala{
		Ministry:     request.Ministry,
    Month:        request.Month,
    Year:         request.Year,
    Data:         request.Data,
	}

	if err := db.Create(&escala).Error; err != nil {
		logger.Errorf("error creating escala: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating escala on database")
		return
	}

	sendSuccess(ctx, "create-escala", escala)
}
