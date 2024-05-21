package handler

import (
	"net/http"

	"github.com/tarcisioads/api_escala/schemas"
	"github.com/gin-gonic/gin"
)

func CreateMemberHandler(ctx *gin.Context) {
	request := CreateMemberRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request body : %v", request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	member := schemas.Member{
    Name: request.Name,
	}

	if err := db.Create(&member).Error; err != nil {
		logger.Errorf("error creating member: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating member on database")
		return
	}

	sendSuccess(ctx, "create-member", member)
}
