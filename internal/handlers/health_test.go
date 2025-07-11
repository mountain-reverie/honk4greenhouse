package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/health", nil)
	c.Request = req
	
	HealthCheck(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.Contains(t, body, `"status":"ok"`)
	assert.Contains(t, body, `"service":"honk4greenhouse"`)
}

func TestHealthCheck_ContentType(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/health", nil)
	req.Header.Set("Accept", "application/json")
	c.Request = req
	
	HealthCheck(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	contentType := w.Header().Get("Content-Type")
	assert.True(t, strings.Contains(contentType, "application/json"))
}