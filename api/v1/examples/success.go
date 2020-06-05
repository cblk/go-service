package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type SuccessResponse struct {
	response.Response
	Data models.Task `json:"data"`
}

func Success(ctx *gin.Context) {
	task := models.NewTask()
	r := &SuccessResponse{}
	r.Data = *task
	r.Success(ctx)
}
