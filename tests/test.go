package tests

import (
	"encoding/json"
	"go-service/forms"
	"go-service/http/app"
	"go-service/utils"
	"net/http"
	"net/http/httptest"
	"strings"
)

type M map[string]interface{}

func (t M) String() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}

func TestIndex() error {
	return utils.Try(func() {
		router := app.GetHttpApplication()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		utils.PanicErr(err)
		router.ServeHTTP(w, req)

		utils.PanicBool(w.Code != http.StatusOK, "test code")
		utils.P(w.Body.String())
	})
}

func TestSendTask() error {
	return utils.Try(func() {
		router := app.GetHttpApplication()
		w := httptest.NewRecorder()

		task := forms.Task{
			Type:  "article",
			Url:   "https://github.com/gin-gonic/gin",
			AppId: "123456789",
		}

		req, err := http.NewRequest("POST", "/api/tasks", strings.NewReader(task.String()))
		utils.PanicErr(err)
		router.ServeHTTP(w, req)

		utils.P(w.Body.String())
		utils.PanicBool(w.Code != http.StatusOK, "test code")
	})
}
