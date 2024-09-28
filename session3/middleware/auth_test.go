package middleware_test

import (
	"belajargolangpart2/session3/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_PositiveCase(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//set router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid, ini mock nya saja
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" dengan token valid, hit ke mock nya
	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "valid-token")

	//buat response recoreder untuk merekam response
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	// Periksa status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Periksa body response
	assert.JSONEq(t, `{"message": "Private data"}`, w.Body.String())
}

func TestAuthMiddleware_NegativeCase_NoToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//set router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid, ini mock nya saja
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" tanpa token, hit ke mock nya
	req, _ := http.NewRequest("GET", "/private", nil)

	//buat response recoreder untuk merekam response
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	//periksa status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	//periksa body response
	assert.Contains(t, w.Body.String(), "Authorization token is required")
}

func TestAuthMiddleware_NegativeCase_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//set router
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// Handler yang hanya dapat diakses dengan token valid, ini mock nya saja
	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Private data"})
	})

	// Buat permintaan HTTP GET ke endpoint "/private" dengan token valid, hit ke mock nya
	req, _ := http.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "invalid-token")

	//buat response recoreder untuk merekam response
	w := httptest.NewRecorder()

	// Lakukan permintaan
	r.ServeHTTP(w, req)

	//periksa status code
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	//periksa body response
	assert.Contains(t, w.Body.String(), "Invalid authorization token")
}
