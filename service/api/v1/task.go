package v1

import (
	"go_service/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"go_service/config"
	"go_service/forms"
	"go_service/library/logy"
	"go_service/models"
	"net/http"
)

func PostTask(ctx *gin.Context) {

	task := &forms.Task{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		logy.ErrorC(ctx, "PostTask", err).Error()
		service.ErrorMsg(ctx, service.CONST_ResultCode_ParseJSON_Error)
		return
	}

	t := models.NewTask()
	if t == nil {
		service.ErrorMsg(ctx, service.CONST_ResultCode_Server_error)
		return
	}
	t.AppId = task.AppId
	t.Type = task.Type
	t.Input = models.M{
		"data": task.Url,
	}.String()

	if err := t.Save(config.GetDB()); err != nil {
		logy.ErrorC(ctx, "PostTask", err).Error()
		service.ErrorMsg(ctx, service.CONST_ResultCode_Server_error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"task_id": t.TaskID,
		},
	})

	service.Success(ctx, http.StatusBadRequest, nil)
}

func GetTask(ctx *gin.Context) {
	_id := ctx.Param("id")
	if _id == "" {
		logy.ErrorC(ctx, "GetTaskParam", errors.New("please input id")).Error()
		service.ErrorMsg(ctx, service.CONST_ResultCode_InputParamter_Empty)
		return
	}

	t := models.NewTask()
	err := t.GetTaskStatus(config.GetDB(), _id)
	if err != nil {
		logy.ErrorC(ctx, "GetTaskStatus", err).Error()
		service.ErrorMsg(ctx, service.CONST_ResultCode_Server_error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": t.Status,
		"data": gin.H{
			"url": t.Output,
		},
	})

	service.Success(ctx, http.StatusBadRequest, nil)
}
