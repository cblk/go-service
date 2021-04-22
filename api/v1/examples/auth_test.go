package examples_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"go_service/api/v1/examples"
	"go_service/tests"
	"go_service/tests/api/v1"

	"github.com/magiconair/properties/assert"
)

func TestAuthNoUsername(t *testing.T) {
	r := callAuthAPI("", "password")
	v1.AssertValidationErrorResponse(t, r, "username", "required")
}

func TestAuthNoPassword(t *testing.T) {
	r := callAuthAPI("admin", "")
	v1.AssertValidationErrorResponse(t, r, "password", "required")
}

func TestAuthWrongUsername(t *testing.T) {
	r := callAuthAPI("admin123", "admin")
	v1.AssertValidationErrorResponse(t, r, "username", "user_not_exist")
}

func TestAuthWrongPassword(t *testing.T) {
	r := callAuthAPI("admin", "admin123")
	v1.AssertValidationErrorResponse(t, r, "password", "incorrect_password")
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

func TestJsonInput(t *testing.T) {
	jsonStr := []byte(`{"username":"admin","password":"admin"}`)
	req, err := http.NewRequest("POST", "/v1/examples/auth", bytes.NewBuffer(jsonStr))

	assert.Equal(t, err, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(len(jsonStr)))

	w := httptest.NewRecorder()
	tests.Application.ServeHTTP(w, req)

	authResponse := &examples.AuthResponse{}

	err = json.Unmarshal(w.Body.Bytes(), authResponse)
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
