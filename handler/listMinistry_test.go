package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
  "encoding/json"
  "github.com/stretchr/testify/assert"
	"github.com/tarcisioads/api_escala/config"
)


func TestListMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpMockRouter()
	basePath := config.GetBasePath() 
  v1 := router.Group(basePath)
  {
    v1.POST("/ministry", CreateMinistryHandler)
    v1.GET("/ministry/:id", UpdateMinistryHandler)
  }
  InitializeHandler()


  t.Run("Should get list ministries", func(t *testing.T) {
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

    reqGet, _ := http.NewRequest("GET", basePath+"/ministries", reqBody)
    wGet := httptest.NewRecorder()
    router.ServeHTTP(wGet, reqGet)
    
    assert.Equal(t, http.StatusOK, wGet.Code)
  })


  t.Run("Should not list a ministries", func(t *testing.T) {
    req, _ := http.NewRequest("GET", basePath+"/ministries", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotFound, w.Code)
  })


}


