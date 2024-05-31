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
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramParameter").Error())
		return
	}
	ministry := schemas.Ministry{}

	if err := db.Model(&schemas.Ministry{}).Preload("Members").First(&ministry, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "ministry not found")
		return
	}

  members := []schemas.Member{}
  for _, name := range request.Members {
    member := schemas.Member{}
    if err := db.Where("name = ?", *name).First(&member).Error; err != nil {
      member.Name = *name
      if err := db.Create(&member).Error; err != nil {
        logger.Errorf("error creating member: %v", err.Error())
        SendError(ctx, http.StatusInternalServerError, "error creating member on database")
        return
      }
    }
    members = append(members, member)
  }

	// Update ministry
	if request.Name != "" {
		ministry.Name = request.Name
	}

  ministry.Members = members

	// Save ministry
	if err := db.Save(&ministry).Error; err != nil {
		logger.Errorf("error updating ministry: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error updating ministry")
		return
	}
	SendSuccess(ctx, "update-ministry", ministry)
}
