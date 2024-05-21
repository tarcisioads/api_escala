package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarcisioads/api_escala/schemas"
)

func UpdateMemberHandler(ctx *gin.Context) {
	request := UpdateMemberRequest{}

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
	member := schemas.Member{}

	if err := db.First(&member, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "member not found")
		return
	}
	// Update member
	if request.Name != "" {
		member.Name = request.Name
	}

	// Save member
	if err := db.Save(&member).Error; err != nil {
		logger.Errorf("error updating member: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating member")
		return
	}
	sendSuccess(ctx, "update-member", member)
}
