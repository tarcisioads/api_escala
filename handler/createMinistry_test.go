package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
  "github.com/stretchr/testify/assert"
	"github.com/tarcisioads/api_escala/config"
)


func TestCreateMinistryHandler(t *testing.T) {
  config.Init(config.TestMode)
	router := SetUpMockRouter()
	basePath := "/api/v1"
  v1 := router.Group(basePath)
  {
    v1.POST("/ministry", CreateMinistryHandler)
  }
  InitializeHandler()


  t.Run("Should create a new ministry", func(t *testing.T) {

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

  });

  t.Run("Should not create a new ministry", func(t *testing.T) {

    ministry := CreateMinistryRequest{
      Name: "",
      Members: nil,
    }
  
    jsonData, _ := json.Marshal(ministry) 
    reqBody := bytes.NewBuffer(jsonData)
  
    req, _ := http.NewRequest("POST", basePath+"/ministry", reqBody)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
  
    assert.Equal(t, http.StatusBadRequest, w.Code)

  })


  }


