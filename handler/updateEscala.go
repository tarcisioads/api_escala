package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateEscalaHandler(ctx *gin.Context) {
	request := UpdateEscalaRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	escala := schemas.Escala{}

	if err := db.First(&escala, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "escala not found")
		return
	}
	// Update escala
	if request.Ministry != "" {
		escala.Ministry = request.Ministry
	}

	if request.Month > 0 {
		escala.Month = request.Month
	}

	if request.Year > 0 {
		escala.Year = request.Year
	}

	if request.Data != "" {
		escala.Data = request.Data
	}

	// Save escala
	if err := db.Save(&escala).Error; err != nil {
		logger.Errorf("error updating escala: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error updating escala")
		return
	}
	SendSuccess(ctx, "update-escala", escala)
}
