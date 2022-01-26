package v1

import (
	"github.com/cblk/go-service/api/v1/examples"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func InitRoutes(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	examplesGroup := v1g.Group("examples", "ResponseExamples", "Example APIs for response format")

	examplesGroup.GET("/success", []fizz.OperationOption{
		fizz.Summary("Get a success response with an example model"),
	}, tonic.Handler(examples.Success, 200))
}
