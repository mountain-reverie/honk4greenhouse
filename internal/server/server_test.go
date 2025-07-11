package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_EndToEnd_HomeHTML(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "text/html")
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), "<!DOCTYPE html>")
	assert.Contains(t, w.Body.String(), "Honk4Greenhouse")
	assert.Contains(t, w.Body.String(), "Hello World! Welcome to honk4greenhouse - your greenhouse design companion")
}

func TestServer_EndToEnd_HomeJSON(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.Contains(t, body, `"message":"Hello World! Welcome to honk4greenhouse - your greenhouse design companion"`)
	assert.Contains(t, body, `"version":"1.0.0"`)
}

func TestServer_EndToEnd_Health(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.Contains(t, body, `"status":"ok"`)
	assert.Contains(t, body, `"service":"honk4greenhouse"`)
}

func TestServer_EndToEnd_NotFound(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestServer_EndToEnd_BrowserRequest(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	
	body := w.Body.String()
	assert.True(t, strings.Contains(body, "<!DOCTYPE html>"))
	assert.True(t, strings.Contains(body, "<title>Honk4Greenhouse</title>"))
}

func TestServer_EndToEnd_APIRequest(t *testing.T) {
	t.Parallel()
	r := New()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
	
	body := w.Body.String()
	assert.True(t, strings.Contains(body, `"message"`))
	assert.True(t, strings.Contains(body, `"version"`))
}