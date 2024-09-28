package router_test

import (
	"belajargolangpart2/session3/router"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter_RootHanlder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//inisiasi router
	r := gin.Default()
	router.SetupRouter(r)

	//buat request
	req, _ := http.NewRequest("GET", "/", nil)

	//buat recod
	w := httptest.NewRecorder()

	//lakukan request
	r.ServeHTTP(w, req)

	//periksa status code
	assert.Equal(t, http.StatusOK, w.Code)

	//periksa body response
	expectedBody := `{"message":"Hallo dari gin"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestPostHandler_PositiveCase(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//inisiasi router
	r := gin.Default()
	router.SetupRouter(r)

	//persiapan data json
	requestBody := map[string]string{"message": "Test message"}
	requestBodyBytes, _ := json.Marshal((requestBody))

	// Buat permintaan HTTP POST dengan data JSON yang valid
	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")

	//buat recorder
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Periksa body response
	expectedBody := `{"message": "Test message"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestPostHandler_NegativeCase_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//inisiasi router
	r := gin.Default()
	router.SetupRouter(r)

	// Buat permintaan HTTP POST dengan data JSON yang tidak valid
	req, _ := http.NewRequest("POST", "/private/post", bytes.NewBufferString("{Invalid JSON}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "valid-token")

	//buat recorder
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Periksa body response
	assert.Contains(t, w.Body.String(), "invalid character")
}

func TestPostHandler_NegativeCase_NoAuthHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//inisiasi router
	r := gin.Default()
	router.SetupRouter(r)

	req, _ := http.NewRequest("POST", "/private/post", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	assert.Contains(t, w.Body.String(), "{\"error\":\"Authorization token is required\"}")
}
