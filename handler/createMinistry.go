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

  members := []schemas.Member{}
  for _, name := range request.Members {
    member := schemas.Member{}
    if err := db.Where("name = ?", *name).First(&member).Error; err != nil {
      member.Name = *name
      if err := db.Create(&member).Error; err != nil {
        logger.Errorf("error creating member: %v", err.Error())
        sendError(ctx, http.StatusInternalServerError, "error creating member on database")
        return
      }
    }
    members = append(members, member)
  }

	ministry := schemas.Ministry{
    Name: request.Name,
    Members: members,
	}

	if err := db.Create(&ministry).Error; err != nil {
		logger.Errorf("error creating ministry: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating ministry on database")
		return
	}

	sendSuccess(ctx, "create-ministry", ministry)
}
