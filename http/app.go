package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-service/http/api"
	"net/http"
)

func GetHttpApplication() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(r *gin.Context) {
		r.String(http.StatusOK, "ok")
	})

	// 服务端API
	apiGroup := r.Group("/api")
	api.InitRouterV1(apiGroup)

	return r
}
