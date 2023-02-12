package apiserver

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleFunc(t *testing.T) {
	mockResponse := `{"message":"pong"}`
	s := New(NewConfig())
	r := s.router
	r.GET("/", s.handleFunc)
	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	responseData, _ := io.ReadAll(rec.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, rec.Code)
}
