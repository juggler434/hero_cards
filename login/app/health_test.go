package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	srv := Server{
		router: http.DefaultServeMux,
	}
	srv.routes()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	srv.router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status: %d, got status:  %d", http.StatusOK, w.Code)
	}
}
