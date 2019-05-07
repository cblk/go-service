package api

import (
	"net/http"

	"go_service/library/logy"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetHttpApplication() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	gin.DefaultWriter = logy.GetLogStdoutInstance()
	gin.DefaultErrorWriter = logy.GetLogStdoutInstance()
	r.Use(gin.LoggerWithWriter(logy.GetLogStdoutInstance()))
	r.Use(gin.RecoveryWithWriter(logy.GetLogStdoutInstance()))

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
