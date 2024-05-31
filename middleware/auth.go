package middleware

import (
	"net/http"
	"os"
	"time"
  "fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tarcisioads/api_escala/initializers"
	"github.com/tarcisioads/api_escala/schemas"
	"github.com/tarcisioads/api_escala/handler"
)


func RequireAuth(ctx *gin.Context) {
  InitializeMiddleware()
  tokenString, err := ctx.Cookie("Authorization")

  if err != nil {
    logger.Errorf("Failed to get Authorization cookie: %v", err.Error())
    handler.SendError(ctx, http.StatusUnauthorized, "Failed to get Authorization cookie")
    return
  } 

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }

    return []byte(os.Getenv("SECRET")), nil
  })

  if err != nil {
    logger.Errorf("Failed to parse token: %v", err.Error())
    handler.SendError(ctx, http.StatusUnauthorized, "Failed to parse token")
    return
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

    if float64(time.Now().Unix()) > claims["exp"].(float64) {
      logger.Errorf("Expired token: %v", err.Error())
      handler.SendError(ctx, http.StatusUnauthorized, "Expired token")
      return
    }

    var user schemas.User
    initializers.DB.First(&user, claims["sub"])

    if user.ID == 0 {
      logger.Errorf("Invalid token: %v", err.Error())
      handler.SendError(ctx, http.StatusUnauthorized, "Invalid token")
      return
    }

    ctx.Set("user", user)

    ctx.Next()
  } else {
    logger.Errorf("Invalid token: %v", err.Error())
    handler.SendError(ctx, http.StatusUnauthorized, "Invalid token")
    return
  }
}

