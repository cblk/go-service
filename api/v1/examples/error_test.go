package examples_test

import (
	"go_service/tests"
	v1 "go_service/tests/api/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/examples/error", nil)

	w := httptest.NewRecorder()
	tests.Application.ServeHTTP(w, req)

	v1.AssertValidationErrorResponse(t, w, "username", "user_not_exist")
}
