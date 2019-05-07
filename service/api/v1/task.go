package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"go_service/config"
	"go_service/forms"
	"go_service/library/logy"
	"go_service/models"
)

func PostTask(ctx *gin.Context) {

	task := &forms.Task{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		logy.Error("PostTask", err)
		ctx.String(http.StatusBadRequest, "参数错误")
		return
	}

	t := models.NewTask()
	t.AppId = task.AppId
	t.Type = task.Type
	t.Input = models.M{
		"data": task.Url,
	}.String()

	if err := t.Save(config.GetDB()); err != nil {
		logy.Error("PostTask", err)
		ctx.String(http.StatusBadRequest, "任务数据保存失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"task_id": t.TaskID,
		},
	})
}

func GetTask(ctx *gin.Context) {
	_id := ctx.Param("id")
	if _id == "" {
		logy.Error("GetTaskParam", errors.New("please input id"))
		ctx.String(http.StatusBadRequest, "please input id")
		return
	}

	t := models.NewTask()
	err := t.GetTaskStatus(config.GetDB(), _id)
	if err != nil {
		logy.Error("GetTaskStatus", err)
		ctx.String(http.StatusBadRequest, "获取任务失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": t.Status,
		"data": gin.H{
			"url": t.Output,
		},
	})
}
