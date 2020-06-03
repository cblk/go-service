package api

import (
	logy "github.com/sirupsen/logrus"
	"go_service/api/v1"
	"go_service/internal/catch"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitRouterV1(r *gin.RouterGroup) {

	v1g := r.Group("v1")

	v1g.GET("tests", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// 提交任务
	v1g.POST("/tasks", v1.PostTask)

	// 得到任务
	v1g.GET("/tasks/:id", v1.GetTask)

	// mock test
	v1g.GET("mock", func(ctx *gin.Context) {
		defer func() {
			revertFunc := func(params ...interface{}) {
				logy.Info("testTx")
			}

			catch.Finally(recover(), revertFunc, "")
			ctx.String(http.StatusInternalServerError, "panic tx")
		}()

		var i int
		var j int
		j = 10
		w := j / i

		logy.Info(strconv.Itoa(w))

		ctx.String(http.StatusOK, "ok")
	})

	responseTestGroup := v1g.Group("response")
	responseTestGroup.GET("/success", v1.Success)
	responseTestGroup.GET("/error", v1.Error)
	responseTestGroup.GET("/exception", v1.Exception)
	responseTestGroup.POST("/auth", v1.Auth)
}
