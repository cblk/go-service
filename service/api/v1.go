package api

import (
	"github.com/gin-gonic/gin"
	"go_service/service/api/v1"
	"net/http"
)

func InitRouterV1(r *gin.RouterGroup) {
	r.GET("tests", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// 提交任务
	r.POST("/tasks", v1.PostTask)

	// 得到任务
	r.GET("/tasks/:id", v1.GetTask)
}
