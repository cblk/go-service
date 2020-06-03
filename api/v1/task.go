package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go_service/api/v1/response"

	logy "github.com/sirupsen/logrus"
	"go_service/config"
	"go_service/forms"
	"go_service/models"
)

func PostTask(ctx *gin.Context) {

	task := &forms.Task{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		logy.Error("PostTask", err)
		response.Error(ctx, "invalid_params", nil)
	}

	t := models.NewTask()
	t.AppId = task.AppId
	t.Type = task.Type
	t.Input = models.M{
		"data": task.Url,
	}.String()

	if err := t.Save(config.GetDB()); err != nil {
		logy.Error("PostTask", err)
		response.Exception(ctx, "database_error")
	}

	response.Success(ctx, gin.H{
		"task_id": t.TaskID,
	})
}

func GetTask(ctx *gin.Context) {
	_id := ctx.Param("id")
	if _id == "" {
		logy.Error("GetTaskParam", errors.New("please input id"))
		response.Error(ctx, "invalid_params", nil)
	}

	t := models.NewTask()
	err := t.GetTaskStatus(config.GetDB(), _id)
	if err != nil {
		logy.Error("GetTaskStatus", err)
		response.Error(ctx, "get_task_failed", nil)
	}

	response.Success(ctx, gin.H{
		"url": t.Output,
	})
}
