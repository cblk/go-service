package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"go_service/config"
	"os"
)

func GetHttpApplication() *gin.Engine {
	cfg := config.GetConfig()
	gin.SetMode(cfg.GetString("gin.mode"))

	engine := gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.LoggerWithWriter(os.Stdout))
	engine.Use(gin.RecoveryWithWriter(os.Stdout))

	fizzEngine := fizz.NewFromEngine(engine)

	// v1 api
	InitRouterV1(fizzEngine)

	// Serve OpenAPI specifications
	infos := &openapi.Info{
		Title:       "Go service",
		Description: "A template for Golang API server",
		Version:     "1.0.0",
	}

	fizzEngine.GET("/openapi.yml", nil, fizzEngine.OpenAPI(infos, "yaml"))

	return engine
}
