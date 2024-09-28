package handler_test

import (
	"belajargolangpart2/session3/handler"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("positive case - correct message", func(t *testing.T) {
		expectedOutput := "Hallo dari gin"
		actualOutput := handler.Gethellomessage()
		require.Equal(t, expectedOutput, actualOutput, "the message should be '%s'", expectedOutput)
	})
}

func TestRootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", handler.RootHandler)

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Hallo dari gin"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

type JsonRequest struct {
	Message string `json:"message"`
}

func TestPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//setup router
	r := gin.Default()
	r.POST("/", handler.PostHandler)

	t.Run("Positive Case", func(t *testing.T) {
		// persiapan data json
		requestBody := JsonRequest{Message: "Hello from test!"}
		requestBodyBytes, _ := json.Marshal(requestBody)

		//buat persiapan http post
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		//buat response recorder
		w := httptest.NewRecorder()

		//lakukan permintaan
		r.ServeHTTP(w, req)

		//periksa status code
		assert.Equal(t, http.StatusOK, w.Code)

		//periksa body response
		expectedBody := `{"message":"Hello from test!"}`
		assert.JSONEq(t, expectedBody, w.Body.String())

	})

	t.Run("Negative case - EOF error", func(t *testing.T) {
		//persiapan data json yang salah
		requestBody := ""
		requestBodyBytes := []byte(requestBody)

		//buat permintaan http post
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Contest-Type", "application/json")

		//buat response recoreder untuk merekam response
		w := httptest.NewRecorder()

		//lakukan permintaan
		r.ServeHTTP(w, req)

		//periksa status code
		assert.Equal(t, http.StatusBadRequest, w.Code)

		//periksa body response
		assert.Contains(t, w.Body.String(), "{\"error\":\"EOF\"}")

	})

}
