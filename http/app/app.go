package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHttpApplication() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(r *gin.Context) {
		r.String(http.StatusOK, "ok")
	})

	// 服务端API
	api := r.Group("/api")
	InitRouterV1(api)

	return r
}
