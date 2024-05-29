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


func TestUpdateMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpMockRouter()
	basePath := config.GetBasePath() 
  v1 := router.Group(basePath)
  {
    v1.POST("/ministry", CreateMinistryHandler)
    v1.PUT("/ministry/:id", UpdateMinistryHandler)
  }
  InitializeHandler()


  t.Run("Should Update a ministry", func(t *testing.T) {
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
    reqUpdate, _ := http.NewRequest("PUT", basePath+"/ministry/"+id, reqBody)
    wUpdate := httptest.NewRecorder()
    router.ServeHTTP(wUpdate, reqUpdate)
    
    assert.Equal(t, http.StatusOK, wUpdate.Code)
  })


  t.Run("Should not update a ministry", func(t *testing.T) {
    id := ""
    req, _ := http.NewRequest("PUT", basePath+"/ministry/"+id, nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotFound, w.Code)
  })


  }


