package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_service/service/api"
	"net/http"
)

func GetHttpApplication() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(r *gin.Context) {
		r.String(http.StatusOK, "ok")
	})

	r.GET("/health", func(r *gin.Context) {
		r.String(http.StatusOK, "ok")
	})

	// 服务端API
	apiGroup := r.Group("/api")
	api.InitRouterV1(apiGroup)

	return r
}
