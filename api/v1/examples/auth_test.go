package examples_test

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"go_service/api/v1/examples"
	"go_service/api/v1/response"
	"go_service/tests"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestAuthNoUsername(t *testing.T) {
	r := callAuthAPI("", "password")
	assert.Equal(t, r.Code, 400)

	// Deserialize response message

	validationResponse := &response.ValidationErrorResponse{}

	err := json.Unmarshal(r.Body.Bytes(), validationResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, validationResponse.GetErrorType(), "validation_error")
	assert.Equal(t, validationResponse.GetFieldName(), "username")
	assert.Equal(t, validationResponse.GetFieldMessage(), "required")
}

func TestAuthNoPassword(t *testing.T) {
	r := callAuthAPI("admin", "")
	assert.Equal(t, r.Code, 400)

	// Deserialize response message

	validationResponse := &response.ValidationErrorResponse{}

	err := json.Unmarshal(r.Body.Bytes(), validationResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, validationResponse.GetErrorType(), "validation_error")
	assert.Equal(t, validationResponse.GetFieldName(), "password")
	assert.Equal(t, validationResponse.GetFieldMessage(), "required")
}

func TestAuthWrongUsername(t *testing.T) {
	r := callAuthAPI("admin123", "admin")
	assert.Equal(t, r.Code, 400)

	// Deserialize response message

	validationResponse := &response.ValidationErrorResponse{}

	err := json.Unmarshal(r.Body.Bytes(), validationResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, validationResponse.GetErrorType(), "validation_error")
	assert.Equal(t, validationResponse.GetFieldName(), "username")
	assert.Equal(t, validationResponse.GetFieldMessage(), "user_not_exist")
}

func TestAuthWrongPassword(t *testing.T) {
	r := callAuthAPI("admin", "admin123")
	assert.Equal(t, r.Code, 400)

	// Deserialize response message

	validationResponse := &response.ValidationErrorResponse{}

	err := json.Unmarshal(r.Body.Bytes(), validationResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, validationResponse.GetErrorType(), "validation_error")
	assert.Equal(t, validationResponse.GetFieldName(), "password")
	assert.Equal(t, validationResponse.GetFieldMessage(), "incorrect_password")
}

func TestAuthSuccess(t *testing.T) {
	r := callAuthAPI("admin", "admin")
	assert.Equal(t, r.Code, 200)

	// Deserialize response message

	authResponse := &examples.AuthResponse{}

	err := json.Unmarshal(r.Body.Bytes(), authResponse)
	assert.Equal(t, err, nil)

	assert.Equal(t, authResponse.GetMessage(), "success")
	assert.Equal(t, authResponse.Data.Username, "admin")
	assert.Equal(t, authResponse.Data.Password, "admin")
}

func callAuthAPI(username, password string) *httptest.ResponseRecorder {

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	req, _ := http.NewRequest("POST", "/v1/examples/auth", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()
	tests.Application.ServeHTTP(w, req)

	return w
}
