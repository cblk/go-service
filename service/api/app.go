package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// v1 api
	InitRouterV1(apiGroup)

	return r
}
