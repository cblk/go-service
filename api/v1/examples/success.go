package examples

import (
	"github.com/gin-gonic/gin"
	"go_service/api/v1/response"
	"go_service/models"
)

type SuccessResponse struct {
	response.Response
	Data models.Task `json:"data" description:"The created task"`
}

func Success(ctx *gin.Context) (*SuccessResponse, error) {
	task := models.NewTask()
	r := &SuccessResponse{}
	r.Data = *task
	return r, nil
}
