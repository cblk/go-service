package api

import (
	"net/http"
	"strconv"

	"go_service/library/catch"
	"go_service/library/logy"
	"go_service/service/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouterV1(r *gin.RouterGroup) {
	r.GET("tests", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// 提交任务
	r.POST("/tasks", v1.PostTask)

	// 得到任务
	r.GET("/tasks/:id", v1.GetTask)

	// mock test
	r.GET("mock", func(ctx *gin.Context) {
		defer func() {
			revertFunc := func(params ...interface{}) {
				logy.Info("testTx", nil)
			}

			catch.Finally(recover(), revertFunc, "")
			ctx.String(http.StatusInternalServerError, "panic tx")
		}()

		var i int
		var j int
		j = 10
		w := j / i

		logy.Info(strconv.Itoa(w), nil)

		ctx.String(http.StatusOK, "ok")
	})
}
