package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHome_HTML_Response(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "text/html")
	c.Request = req
	
	Home(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), "<!DOCTYPE html>")
	assert.Contains(t, w.Body.String(), "Honk4Greenhouse")
	assert.Contains(t, w.Body.String(), "Hello World! Welcome to honk4greenhouse - your greenhouse design companion")
}

func TestHome_HTML_Response_EmptyAccept(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/", nil)
	c.Request = req
	
	Home(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), "<!DOCTYPE html>")
}

func TestHome_JSON_Response(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	c.Request = req
	
	Home(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.Contains(t, body, `"message":"Hello World! Welcome to honk4greenhouse - your greenhouse design companion"`)
	assert.Contains(t, body, `"version":"1.0.0"`)
}

func TestHome_JSON_Response_APIClient(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "curl/7.68.0")
	c.Request = req
	
	Home(c)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.True(t, strings.Contains(body, `"message"`))
	assert.True(t, strings.Contains(body, `"version"`))
}

func BenchmarkHome_HTML(b *testing.B) {
	gin.SetMode(gin.TestMode)
	
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "text/html")
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for b.Loop() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		
		Home(c)
	}
}

func BenchmarkHome_JSON(b *testing.B) {
	gin.SetMode(gin.TestMode)
	
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for b.Loop() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		
		Home(c)
	}
}