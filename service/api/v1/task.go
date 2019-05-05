package v1

import (
	"github.com/gin-gonic/gin"
	"go-service/config"
	"go-service/forms"
	"go-service/models"
	"go-service/utils"
	"net/http"
)

func PostTask(ctx *gin.Context) {
	if err := utils.Try(func() {
		task := &forms.Task{}
		utils.PanicWrap(ctx.ShouldBindJSON(task), "参数错误")

		t := models.NewTask()
		t.AppId = task.AppId
		t.Type = task.Type
		t.Input = models.M{
			"data": task.Url,
		}.String()

		utils.PanicWrap(t.Save(config.GetDB()), "任务数据保存失败")

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

func GetTask(ctx *gin.Context) {
	if err := utils.Try(func() {
		_id := ctx.Param("id")
		utils.PanicBool(_id == "", "please input id")

		t := models.NewTask()
		utils.PanicWrap(t.GetTaskStatus(config.GetDB(), _id), "获取任务失败")

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
