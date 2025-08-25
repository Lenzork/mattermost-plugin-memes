package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestServeTemplateJPEG_MultipleTextParams(t *testing.T) {
	// Set up a request with multiple text parameters
	req := httptest.NewRequest("GET", "/templates/brace-yourselves.jpg?text=top+line&text=bottom+line", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we get a successful response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))
	assert.Equal(t, "public, max-age=604800", rec.Header().Get("Cache-Control"))

	// Check that we have a non-empty JPEG body
	body := rec.Body.Bytes()
	assert.NotEmpty(t, body)

	// Check that it's a valid JPEG (starts with JPEG magic bytes)
	assert.True(t, len(body) > 2)
	assert.Equal(t, byte(0xFF), body[0]) // JPEG magic bytes
	assert.Equal(t, byte(0xD8), body[1])
}

func TestServeTemplateJPEG_SingleTextParam(t *testing.T) {
	// Set up a request with a single text parameter
	req := httptest.NewRequest("GET", "/templates/brace-yourselves.jpg?text=single+line", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we get a successful response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))

	// Check that we have a non-empty JPEG body
	body := rec.Body.Bytes()
	assert.NotEmpty(t, body)
	assert.True(t, len(body) > 2)
	assert.Equal(t, byte(0xFF), body[0]) // JPEG magic bytes
	assert.Equal(t, byte(0xD8), body[1])
}

func TestServeTemplateJPEG_BackwardsCompatibility_TParam(t *testing.T) {
	// Set up a request with old 't' parameter (backwards compatibility)
	req := httptest.NewRequest("GET", "/templates/brace-yourselves.jpg?t=legacy+text", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we get a successful response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))

	// Check that we have a non-empty JPEG body
	body := rec.Body.Bytes()
	assert.NotEmpty(t, body)
	assert.True(t, len(body) > 2)
	assert.Equal(t, byte(0xFF), body[0]) // JPEG magic bytes
	assert.Equal(t, byte(0xD8), body[1])
}

func TestServeTemplateJPEG_TextOverridesTParam(t *testing.T) {
	// Set up a request with both 'text' and 't' parameters - 'text' should take precedence
	req := httptest.NewRequest("GET", "/templates/brace-yourselves.jpg?text=new+text&t=old+text", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we get a successful response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))

	// Check that we have a non-empty JPEG body
	body := rec.Body.Bytes()
	assert.NotEmpty(t, body)
}

func TestServeTemplateJPEG_InvalidTemplate(t *testing.T) {
	// Set up a request for a non-existent template
	req := httptest.NewRequest("GET", "/templates/non-existent.jpg?text=test", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we get a 404 response
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestServeTemplateJPEG_NoTextParams(t *testing.T) {
	// Set up a request with no text parameters
	req := httptest.NewRequest("GET", "/templates/brace-yourselves.jpg", nil)
	rec := httptest.NewRecorder()

	// Set up the router with the template name parameter
	router := mux.NewRouter()
	router.HandleFunc("/templates/{name}.jpg", serveTemplateJPEG).Methods("GET")

	// Serve the request
	router.ServeHTTP(rec, req)

	// Check that we still get a successful response (empty text slots)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))

	// Check that we have a non-empty JPEG body (the base template image)
	body := rec.Body.Bytes()
	assert.NotEmpty(t, body)
}