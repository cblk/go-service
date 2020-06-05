package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Response
	Data error `json:"data"`
}

type ValidationErrorResponse struct {
	Response

	Data struct {
		FieldName string `json:"field_name"`
		Message   string `json:"message"`
	} `json:"data"`
}

func (r *Response) Success() {
	r.Message = "success"
}

func (r *Response) Error(ctx *gin.Context, errorType string) {
	r.Message = errorType
	ctx.JSON(http.StatusOK, r)
}

func (r *Response) Exception(ctx *gin.Context) {
	r.Message = "internal_server_error"
	ctx.JSON(http.StatusInternalServerError, r)
}

func Error(ctx *gin.Context, errorType string, err error) {
	errorResponse := &ErrorResponse{}
	errorResponse.Data = err
	errorResponse.Error(ctx, errorType)
}

func Exception(ctx *gin.Context, message string) {
	r := &Response{}
	r.Message = message
	r.Exception(ctx)
}
