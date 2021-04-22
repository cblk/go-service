package examples_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_service/api/v1/examples"
	"go_service/tests"

	"github.com/magiconair/properties/assert"
)

func TestSuccessResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/examples/success", nil)

	w := httptest.NewRecorder()
	tests.Application.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)

	successResponse := &examples.SuccessResponse{}

	err := json.Unmarshal(w.Body.Bytes(), successResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, successResponse.GetMessage(), "success")
}
