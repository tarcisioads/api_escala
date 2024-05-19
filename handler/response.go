package handler

import (
	"fmt"
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}

type DeleteEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
type ShowEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
type ListEscalaResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.EscalaResponse `json:"data"`
}
type UpdateEscalaResponse struct {
	Message string                 `json:"message"`
	Data    schemas.EscalaResponse `json:"data"`
}
