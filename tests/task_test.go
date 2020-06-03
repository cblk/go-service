package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_service/api"
	"go_service/utils"
)

func TestIndex(t *testing.T) {
	if err := utils.Try(func() {
		router := api.GetHttpApplication()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		utils.PanicErr(err)

		router.ServeHTTP(w, req)

		utils.P(w.Body.String())
		utils.PanicBool(w.Code != http.StatusOK, "test code")
	}); err != nil {
		t.Fatal(err.Error())
	}
}
