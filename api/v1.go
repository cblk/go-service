package api

import (
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"go_service/api/v1/examples"
	"go_service/api/v1/response"
)

func InitRouterV1(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	examplesGroup := v1g.Group("examples", "ResponseExamples", "Example APIs for response format")

	examplesGroup.GET("/success", []fizz.OperationOption{
		fizz.Summary("Get a success response with an example model"),
	}, tonic.Handler(examples.Success, 200))

	examplesGroup.GET("/error", []fizz.OperationOption{
		fizz.Summary("Get an example error response with validation errors"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(examples.Error, 200))

	examplesGroup.GET("/exception", []fizz.OperationOption{
		fizz.Summary("Get an example exception response"),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil),
	}, tonic.Handler(examples.Exception, 200))

	examplesGroup.POST("/auth", []fizz.OperationOption{
		fizz.Summary("Authentication through password"),
		fizz.Response("400", "exception", response.ValidationErrorResponse{}, nil),
	}, tonic.Handler(examples.Auth, 200))
}
