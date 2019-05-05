package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"portal/internal/cnst"
	"portal/utils"
	"strings"
	"time"
)

func NewTask() *Task {
	var _uuid string
	for {
		if _u, err := uuid.NewUUID(); err != nil {
			utils.Log(err)
		} else {
			_uuid = _u.String()
			break
		}
	}

	return &Task{
		CreatedAt: uint(time.Now().Unix()),
		Status:    cnst.Status.Pending,
		Type:      cnst.TaskType.Article,
		AppId:     "123456789",
		TaskID:    strings.ReplaceAll(_uuid, "-", ""),
		Priority:  5,
		RetryNum:  5,
		Version:   cnst.Version[len(cnst.Version)-1].Db,
	}
}

type Task struct {
	ID         uint   `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt  uint   `json:"created_at,omitempty" db:"created_at" gorm:"index;not null"`                    // 每级任务生成时间
	FinishedAt uint   `json:"finished_at,omitempty" db:"finished_at" gorm:"index;not null"`                  // 每级任务完成时间
	Status     string `json:"status,omitempty" db:"status" gorm:"type:varchar(20);index;not null"`           // 任务状态
	ErrType    uint   `json:"err_type,omitempty" db:"err_type" gorm:"index;not null"`                        // 任务错误类型(404，500等)
	Type       string `json:"type,omitempty" db:"type:varchar(20);type;not null"`                            // 任务类型(article, image)
	AppId      string `json:"app_id,omitempty" db:"app_id" gorm:"type:varchar(100);index;not null"`          // APPID
	TaskID     string `json:"task_id,omitempty" db:"task_id" gorm:"type:varchar(100);unique_index;not null"` // 主任务 ID  (uuid.V4().hex())
	Priority   uint8  `json:"priority,omitempty" db:"priority" gorm:"not null"`                              // 任务优先度 1-9
	Input      string `json:"input,omitempty" db:"input" gorm:"type:text;not null"`                          // 任务参数
	Output     string `json:"output,omitempty" db:"output" gorm:"type:text;not null"`                        // 任务参数
	RetryNum   int    `json:"retry_num,omitempty" db:"retry_num" gorm:"not null"`                            // 任务重试次数
	Version    string `json:"version,omitempty" db:"version" gorm:"type:varchar(20);not null"`               // 任务版本
}

func (t *Task) Save(db *gorm.DB) error {
	return utils.Try(func() {
		t.CreatedAt = uint(time.Now().Unix())
		t.Version = cnst.Version[len(cnst.Version)-1].Db
		utils.PanicErr(db.Create(t).Error)
	})
}

func (t *Task) UpdateStatus(db *gorm.DB, taskId, errType, url, status string) error {
	return utils.Try(func() {
		utils.PanicErr(db.Model(t).Where("task_id = ?", taskId).Updates(
			M{"finished_at": time.Now().Unix(), "output": url, "status": status, "err_type": errType}).Error)
	})
}

func (t *Task) Page(db *gorm.DB, taskId, url, status string) error {
	return utils.Try(func() {
		utils.PanicErr(db.Model(t).Where("task_id = ?", taskId).Updates(
			M{"finished_at": time.Now().Unix(), "output": url, "status": status}).Error)
	})
}

func (t *Task) GetTask(db *gorm.DB, taskId string) error {
	return utils.Try(func() {
		utils.PanicErr(db.Where("task_id = ?", taskId).Find(t).Error)
	})
}

func (t *Task) GetTaskStatus(db *gorm.DB, taskId string) error {
	return utils.Try(func() {
		utils.PanicErr(db.Where("task_id = ?", taskId).Select("status, output").Find(t).Error)
	})
}

func (t *Task) Encode() []byte {
	_dt, err := json.Marshal(t)
	utils.PanicErr(err)
	return _dt
}
