package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func ListEscalasHandler(ctx *gin.Context) {
	escalas := []schemas.Escala{}

	if err := db.Find(&escalas).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing escalas")
		return
	}

	SendSuccess(ctx, "list-escalas", escalas)
}
