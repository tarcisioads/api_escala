package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
  "github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/tarcisioads/api_escala/config"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
  gin.SetMode(gin.TestMode) 
	return router
}

func TestCreateMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpRouter()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.POST("/ministry", CreateMinistryHandler)
	}
  InitializeHandler()

  members := []string{"John"}

  ministry := CreateMinistryRequest{
		Name: "Youth Ministry",
		Members: []*string{&members[0]},
	}

  jsonData, _ := json.Marshal(ministry) 
  reqBody := bytes.NewBuffer(jsonData)

	req, _ := http.NewRequest("POST", basePath+"/ministry", reqBody)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)

}


