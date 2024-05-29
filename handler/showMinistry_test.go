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


func TestShowMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpMockRouter()
	basePath := config.GetBasePath() 
  v1 := router.Group(basePath)
  {
    v1.POST("/ministry", CreateMinistryHandler)
    v1.GET("/ministry/:id", UpdateMinistryHandler)
  }
  InitializeHandler()


  t.Run("Should get a ministry", func(t *testing.T) {
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

    ministry.Name = "Youth Ministry Updated"
    jsonData, _ = json.Marshal(ministry) 
    reqBody = bytes.NewBuffer(jsonData)

    id := "1"
    reqGet, _ := http.NewRequest("GET", basePath+"/ministry/"+id, reqBody)
    wGet := httptest.NewRecorder()
    router.ServeHTTP(wGet, reqGet)
    
    assert.Equal(t, http.StatusOK, wGet.Code)
  })


  t.Run("Should not update a ministry", func(t *testing.T) {
    id := ""
    req, _ := http.NewRequest("GET", basePath+"/ministry/"+id, nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotFound, w.Code)
  })


}


