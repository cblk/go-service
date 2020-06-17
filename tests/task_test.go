package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_service/api"
)

func TestIndex(t *testing.T) {
	router := api.GetHttpApplication()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
		return
	}
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Error("code err")
	}
}
