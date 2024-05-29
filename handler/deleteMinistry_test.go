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


func TestDeleteMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpMockRouter()
	basePath := config.GetBasePath() 
  v1 := router.Group(basePath)
  {
    v1.POST("/ministry", CreateMinistryHandler)
    v1.DELETE("/ministry/:id", DeleteMinistryHandler)
  }
  InitializeHandler()


  t.Run("Should delete a ministry", func(t *testing.T) {
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

    id := "1"
    reqDelete, _ := http.NewRequest("DELETE", basePath+"/ministry/"+id, nil)
    wDelete := httptest.NewRecorder()
    router.ServeHTTP(wDelete, reqDelete)
    assert.Equal(t, http.StatusOK, w.Code)
  })


  t.Run("Should not delete a ministry", func(t *testing.T) {
    id := ""
    req, _ := http.NewRequest("DELETE", basePath+"/ministry/"+id, nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotFound, w.Code)
  })


  }


