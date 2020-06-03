package api

import (
	"go_service/config"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetHttpApplication() *gin.Engine {
	cfg := config.GetConfig()
	gin.SetMode(cfg.GetString("gin.mode"))

	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.LoggerWithWriter(os.Stdout))
	r.Use(gin.RecoveryWithWriter(os.Stdout))

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
