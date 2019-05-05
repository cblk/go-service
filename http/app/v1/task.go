package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-service/forms"
	"go-service/models"
	"go-service/utils"
)

func Post_task(ctx *gin.Context) {
	if err := utils.Try(func() {
		task := &forms.Task{}
		utils.PanicWrap(ctx.ShouldBindJSON(task), "参数错误")

		t := models.NewTask()
		t.AppId = task.AppId
		t.Type = task.Type
		t.Input = models.M{
			"data": task.Url,
		}.String()

		utils.PanicWrap(t.Save(cfg.GetDb()), "任务数据保存失败")

		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"task_id": t.TaskID,
			},
		})

	}); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

}

func Get_task(ctx *gin.Context) {
	if err := utils.Try(func() {
		_id := ctx.Param("id")
		utils.PanicBool(_id == "", "please input id")

		cfg := config.DefaultConfig()

		t := models.NewTask()
		utils.PanicWrap(t.GetTaskStatus(cfg.GetDb(), _id), "获取任务失败")

		ctx.JSON(http.StatusOK, gin.H{
			"code":   0,
			"status": t.Status,
			"data": gin.H{
				"url": t.Output,
			},
		})

	}); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
}
