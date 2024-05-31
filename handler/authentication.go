package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tarcisioads/api_escala/schemas"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignupHandler(ctx *gin.Context) {
	request := SignupUserRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request body : %v", request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Email:    request.Email,
		Password: request.Email,
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		logger.Errorf("Failed to hash password: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user.Password = string(hash)

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}

	SendSuccess(ctx, "create-user", user)
}

func SigninHandler(ctx *gin.Context) {
	request := SigninUserRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request body : %v", request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	if err := db.Model(&schemas.User{}).First(&user, "email = ?", request.Email).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		logger.Errorf("hash password is different error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		logger.Errorf("error signing token: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error signing token")
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)

	SendSuccess(ctx, "signin-user", user)

}


func ValidateHandler(ctx *gin.Context) {
  _, err := ctx.Cookie("token")
  if err != nil {
    ctx.AbortWithStatus(http.StatusUnauthorized)
    return
  }
  ctx.Next()
}
