package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouterV1(r *gin.RouterGroup) {
	r.GET("tests", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// 提交任务
	r.POST("/tasks", v1.Post_task)

	// 得到任务
	r.GET("/tasks/:id", v1.Get_task)

	// 更新任务状态
	r.PUT("/tasks/:id/status", v1.Put_task_status)

}
