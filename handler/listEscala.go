package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ListEscalasHandler(ctx *gin.Context) {
	escalas := []schemas.Escala{}

	if err := db.Find(&escalas).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing escalas")
		return
	}

	sendSuccess(ctx, "list-escalas", escalas)
}
