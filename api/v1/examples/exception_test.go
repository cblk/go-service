package examples_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_service/tests"
	v1 "go_service/tests/api/v1"
)

func TestExceptionResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/examples/exception", nil)

	w := httptest.NewRecorder()
	tests.Application.ServeHTTP(w, req)

	v1.AssertExceptionResponse(t, w, "internal server error")
}
