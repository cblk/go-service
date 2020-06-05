package api

import (
	"github.com/wI2L/fizz"
	"go_service/api/v1/examples"
)

func InitRouterV1(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	examplesGroup := v1g.Group("examples", "ResponseExamples", "Example APIs for response format")
	examplesGroup.GET("/success", nil, examples.Success)
	examplesGroup.GET("/error", nil, examples.Error)
	examplesGroup.GET("/exception", nil, examples.Exception)
	examplesGroup.POST("/auth", nil, examples.Auth)
}
